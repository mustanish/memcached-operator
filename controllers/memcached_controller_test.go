package controllers

import (
	"path/filepath"

	"github.com/example/memcached-operator/api/v1alpha1"
	. "github.com/example/memcached-operator/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Operator Controller", func() {
	var defaultVal *v1alpha1.Memcached
	var expectedVal *v1alpha1.OperatorConstant

	BeforeEach(func() {
		defaultVal = new(v1alpha1.Memcached)
		_ = ParseYAML(filepath.Join("..", "config", "samples", "cache_v1alpha1_memcached.yaml"), defaultVal)
		expectedVal = new(v1alpha1.OperatorConstant)
		_ = ParseYAML(filepath.Join("..", "config", "samples", "openshift_constant.yaml"), expectedVal)
	})

	Context("When both YAML(operator and constant) is expected to have the correct values", func() {
		It("Should match and pass", func() {
			Expect(defaultVal.Spec.Constant).To(Equal(*expectedVal))
		})
	})

	AfterEach(func() {
		defaultVal, expectedVal = nil, nil
	})
})
