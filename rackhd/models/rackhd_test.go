package models_test

import (
	"github.com/RackHD/neighborhood-manager/rackhd/models"
	"github.com/hashicorp/consul/api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RackHD Models", func() {
	var rhd1UUID string
	var rhd1HTTP string
	var rhd1AMQP string
	var rhd2UUID string
	var rhd2HTTP string
	var rhd2AMQP string
	var rhd3UUID string
	var rhd3HTTP string
	var rhd3AMQP string
	var err error
	var rhd *models.RackHD

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	kv := client.KV()

	BeforeEach(func() {
		models.InitBackend()
	})

	AfterEach(func() {
	})

	Describe("Creating RackHD Structs", func() {
		BeforeEach(func() {
			rhd1UUID = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
			rhd1HTTP = "http://10.10.10.10:2020"
			rhd1AMQP = "amqp://localhost/"
		})

		JustBeforeEach(func() {
			rhd, err = models.NewRhd(rhd1UUID, rhd1HTTP, rhd1AMQP)
		})

		Context("when valid data is provided", func() {
			It("should return a RackHD struct", func() {
				Expect(rhd).Should(BeAssignableToTypeOf(&models.RackHD{}))
			})

			It("should populate the fields correctly", func() {
				Expect(rhd.HTTPConf).Should(BeAssignableToTypeOf(models.HTTPConfig{}))
				Expect(rhd.AmqpConf).Should(BeAssignableToTypeOf(models.AmqpConfig{}))
				Expect(rhd.ID).To(Equal(rhd1UUID))
				Expect(rhd.HTTPConf.URL.String()).To(Equal(rhd1HTTP))
				Expect(rhd.AmqpConf.URI.String()).To(Equal(rhd1AMQP))
			})

			It("should not have errored", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when invalid http data is provided", func() {
			BeforeEach(func() {
				rhd1UUID = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
				rhd1HTTP = "^$%#$#$"
				rhd1AMQP = "amqp://localhost/"
			})

			It("should return a partial RackHD struct", func() {
				Expect(rhd).Should(BeAssignableToTypeOf(&models.RackHD{}))
			})

			It("should have errored", func() {
				Expect(err).To(HaveOccurred())
			})

		})

		Context("when invalid amqp data is provided", func() {
			BeforeEach(func() {
				rhd1UUID = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
				rhd1HTTP = "http://10.10.10.10:2020"
				rhd1AMQP = "312"
			})

			It("should return a partial RackHD struct", func() {
				Expect(rhd).Should(BeAssignableToTypeOf(&models.RackHD{}))
			})

			It("should populate the HTTP field correctly", func() {
				Expect(rhd.HTTPConf).Should(BeAssignableToTypeOf(models.HTTPConfig{}))
				Expect(rhd.AmqpConf).Should(BeAssignableToTypeOf(models.AmqpConfig{}))
				Expect(rhd.HTTPConf.URL.String()).To(Equal(rhd1HTTP))
				Expect(rhd.AmqpConf.URI.String()).NotTo(Equal(rhd1AMQP))
			})

			It("should have errored", func() {
				Expect(err).To(HaveOccurred())
			})

		})
	})

	Describe("Creating RackHD Instances", func() {
		BeforeEach(func() {
			rhd1UUID = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
			rhd1HTTP = "http://10.10.10.10:2020"
			rhd1AMQP = "amqp://localhost/"
		})

		AfterEach(func() {
			kv.DeleteTree("rhdman/rhd/"+rhd1UUID, nil)
		})

		JustBeforeEach(func() {
			rhd, err = models.NewRhd(rhd1UUID, rhd1HTTP, rhd1AMQP)
			if err != nil {
				Fail("Could not create object")
			}
			err = models.CreateRhd(rhd)
		})

		Context("when valid data is provided", func() {
			It("should create an instance without error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should populate the DB with the correct values", func() {
				pair, _, err := kv.Get("rhdman/rhd/"+rhd1UUID+"/"+"httpConf", nil)
				if err != nil {
					Fail("Unable to get value from the DB for testing")
				}
				Expect(string(pair.Value)).To(Equal(rhd1HTTP))
				pair, _, err = kv.Get("rhdman/rhd/"+rhd1UUID+"/"+"amqpConf", nil)
				if err != nil {
					Fail("Unable to get value from the DB for testing")
				}
				Expect(string(pair.Value)).To(Equal(rhd1AMQP))
			})
		})
	})

	Describe("Retrieving RackHD Instances", func() {
		BeforeEach(func() {
			rhd1UUID = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
			rhd1HTTP = "http://10.10.10.10:2020"
			rhd1AMQP = "amqp://localhost/"
			rhd1, _ := models.NewRhd(rhd1UUID, rhd1HTTP, rhd1AMQP)
			_ = models.CreateRhd(rhd1)
			rhd2UUID = "2e650685-bbe7-44f6-9c66-bd466d1bc1ab"
			rhd2HTTP = "http://20.20.20.20:4020"
			rhd2AMQP = "amqp://localhost/"
			rhd2, _ := models.NewRhd(rhd2UUID, rhd2HTTP, rhd2AMQP)
			_ = models.CreateRhd(rhd2)
			rhd3UUID = "6c04c204-5c5c-4daf-9d68-7e3f2c76de2d"
			rhd3HTTP = "http://30.30.30.30:6020"
			rhd3AMQP = "amqp://localhost/"
			rhd3, _ := models.NewRhd(rhd3UUID, rhd3HTTP, rhd3AMQP)
			_ = models.CreateRhd(rhd3)

		})

		AfterEach(func() {
			kv.DeleteTree("rhdman/rhd/"+rhd1UUID, nil)
			kv.DeleteTree("rhdman/rhd/"+rhd2UUID, nil)
			kv.DeleteTree("rhdman/rhd/"+rhd3UUID, nil)
		})

		Context("Getting All RackHD Instances", func() {

			It("should not have errored", func() {
				_, err := models.GetAllRhd()
				Expect(err).ToNot(HaveOccurred())
			})

			It("should get all instances in the DB", func() {
				rhds, _ := models.GetAllRhd()
				Expect(len(rhds)).To(Equal(3))
			})
		})

		Context("Getting RackHD by ID", func() {

			It("should not have errored", func() {
				_, err := models.GetRhdByID(rhd2UUID)
				Expect(err).ToNot(HaveOccurred())
			})

			It("should return the right RHD", func() {
				rhd, _ := models.GetRhdByID(rhd2UUID)
				Expect(rhd.ID).To(Equal(rhd2UUID))
				Expect(rhd.HTTPConf.URL.String()).To(Equal(rhd2HTTP))
				Expect(rhd.AmqpConf.URI.String()).To(Equal(rhd2AMQP))
			})
		})

		Context("Getting multiple RackHDs by multiple IDs", func() {

			It("should not have errored", func() {
				arr := []string{rhd2UUID, rhd3UUID}
				_, err := models.GetRhdsByIDs(arr)
				Expect(err).ToNot(HaveOccurred())
			})

			It("should return only 2 RHDs", func() {
				arr := []string{rhd2UUID, rhd3UUID}
				rhds, _ := models.GetRhdsByIDs(arr)
				Expect(len(rhds)).To(Equal(2))
			})
		})
	})

	Describe("Deleting RackHD Instances", func() {
		BeforeEach(func() {
			rhd1UUID = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
			rhd1HTTP = "http://10.10.10.10:2020"
			rhd1AMQP = "amqp://localhost/"
			rhd1, _ := models.NewRhd(rhd1UUID, rhd1HTTP, rhd1AMQP)
			_ = models.CreateRhd(rhd1)
			rhd2UUID = "2e650685-bbe7-44f6-9c66-bd466d1bc1ab"
			rhd2HTTP = "http://20.20.20.20:4020"
			rhd2AMQP = "amqp://localhost/"
			rhd2, _ := models.NewRhd(rhd2UUID, rhd2HTTP, rhd2AMQP)
			_ = models.CreateRhd(rhd2)
			rhd3UUID = "6c04c204-5c5c-4daf-9d68-7e3f2c76de2d"
			rhd3HTTP = "http://30.30.30.30:6020"
			rhd3AMQP = "amqp://localhost/"
			rhd3, _ := models.NewRhd(rhd3UUID, rhd3HTTP, rhd3AMQP)
			_ = models.CreateRhd(rhd3)
		})

		AfterEach(func() {
			kv.DeleteTree("rhdman/rhd/"+rhd1UUID, nil)
			kv.DeleteTree("rhdman/rhd/"+rhd2UUID, nil)
			kv.DeleteTree("rhdman/rhd/"+rhd3UUID, nil)
		})

		Context("Deleting RackHD Instance by ID", func() {
			It("should not have errored", func() {
				err = models.DeleteRhdByID(rhd1UUID)
				Expect(err).ToNot(HaveOccurred())
			})

			It("should have removed the instance", func() {
				_ = models.DeleteRhdByID(rhd1UUID)
				rhds, _ := models.GetAllRhd()
				Expect(len(rhds)).To(Equal(2))
			})
		})

		Context("Deleting multiple RackHDs by multiple IDs", func() {
			It("should not have errored", func() {
				arr := []string{rhd2UUID, rhd3UUID}
				err := models.DeleteRhdsByIDs(arr)
				Expect(err).ToNot(HaveOccurred())
			})

			It("should have removed multiple instances", func() {
				arr := []string{rhd2UUID, rhd3UUID}
				_ = models.DeleteRhdsByIDs(arr)
				rhds, _ := models.GetAllRhd()
				Expect(len(rhds)).To(Equal(1))
			})
		})
	})
})
