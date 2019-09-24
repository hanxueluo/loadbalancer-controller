/*
Copyright 2017 Caicloud authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nginx

import (
	"fmt"
	"reflect"

	lbapi "github.com/caicloud/clientset/pkg/apis/loadbalance/v1alpha2"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	log "k8s.io/klog"
)

var (
	defaultConfig = map[string]string{
		"enable-sticky-sessions": "true",
		"ssl-redirect":           "false",
		"force-ssl-redirect":     "false",
		"enable-vts-status":      "true",
		"proxy-body-size":        "5G",
		"server-tokens":          "false",
		"skip-access-log-urls":   "/nginx_status/format/json",
	}
)

func merge(dst, src map[string]string) map[string]string {
	ret := make(map[string]string)

	for k, v := range dst {
		ret[k] = v
	}
	for k, v := range src {
		ret[k] = v
	}

	return ret
}

func (f *nginx) ensureConfigMaps(lb *lbapi.LoadBalancer) error {
	labels := f.selector(lb)

	defaultLBConfigMapName := "default-lb-config"

	cm, err := f.ensureConfigMap(defaultLBConfigMapName, lb.Namespace, labels, nil)
	if err != nil {
		return err
	}
	if len(cm.Data) == 0 {
		cm, err = f.ensureConfigMap(defaultLBConfigMapName, lb.Namespace, labels, defaultConfig)
		if err != nil {
			return err
		}
	}

	config := merge(cm.Data, lb.Spec.Proxy.Config)
	cmName := fmt.Sprintf(configMapName, lb.Name)
	_, err = f.ensureConfigMap(cmName, lb.Namespace, labels, config)
	if err != nil {
		return err
	}
	tcpcmName := fmt.Sprintf(tcpConfigMapName, lb.Name)
	_, err = f.ensureConfigMap(tcpcmName, lb.Namespace, labels, nil)
	if err != nil {
		return err
	}
	udpcmName := fmt.Sprintf(udpConfigMapName, lb.Name)
	_, err = f.ensureConfigMap(udpcmName, lb.Namespace, labels, nil)

	return err
}

func (f *nginx) ensureConfigMap(name, namespace string, labels, data map[string]string) (*v1.ConfigMap, error) {
	cm, err := f.client.CoreV1().ConfigMaps(namespace).Get(name, metav1.GetOptions{})

	if err != nil && !errors.IsNotFound(err) {
		return nil, err
	}

	if errors.IsNotFound(err) {
		cm = &v1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name:   name,
				Labels: labels,
			},
			Data: data,
		}
		log.Infof("About to craete ConfigMap %v/%v for proxy", namespace, cm.Name)
		_, nerr := f.client.CoreV1().ConfigMaps(namespace).Create(cm)
		if nerr != nil {
			return nil, nerr
		}
	}

	if data == nil {
		// do not update data if data == nil
		// tcp and udp config map will be changed by others
		// the controller only need to create it
		return cm, nil
	}

	if reflect.DeepEqual(cm.Data, data) {
		return cm, nil
	}

	// replace cm.Data of data
	// the data follows the priority
	// 1. lb.Spec.Proxy.Config
	// 2. default config
	cm.Data = data
	log.Infof("About to update ConfigMap %v/%v data", namespace, cm.Name)
	cm, err = f.client.CoreV1().ConfigMaps(namespace).Update(cm)

	return cm, err
}
