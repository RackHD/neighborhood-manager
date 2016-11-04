package models

import (
	"fmt"
	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/boltdb"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
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
	var rhd *RackHD

	wd, _ := os.Getwd()
	testDB := wd + "/test.db"

	BeforeEach(func() {
		boltdb.Register()
		var err error
		db, err = libkv.NewStore(
			store.BOLTDB,
			[]string{testDB},
			&store.Config{
				PersistConnection: true,
				Bucket:            "test",
			},
		)
		if err != nil {
			fmt.Printf("Failed to init DB!\n")
		}
	})

	AfterEach(func() {
		db.Close()
		_ = os.Remove(testDB)
	})

	Describe("Creating RackHD Structs", func() {
		var err error
		BeforeEach(func() {
			rhd1UUID = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
			rhd1HTTP = "http://10.10.10.10:2020"
			rhd1AMQP = "amqp://localhost/"
		})

		JustBeforeEach(func() {
			rhd, err = NewRhd(rhd1UUID, rhd1HTTP, rhd1AMQP)
		})

		Context("when valid data is provided", func() {
			It("should return a RackHD struct", func() {
				Expect(rhd).Should(BeAssignableToTypeOf(&RackHD{}))
			})

			It("should populate the fields correctly", func() {
				Expect(rhd.HTTPConf).Should(BeAssignableToTypeOf(HTTPConfig{}))
				Expect(rhd.AmqpConf).Should(BeAssignableToTypeOf(AmqpConfig{}))
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
				Expect(rhd).Should(BeAssignableToTypeOf(&RackHD{}))
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
				Expect(rhd).Should(BeAssignableToTypeOf(&RackHD{}))
			})

			It("should populate the HTTP field correctly", func() {
				Expect(rhd.HTTPConf).Should(BeAssignableToTypeOf(HTTPConfig{}))
				Expect(rhd.AmqpConf).Should(BeAssignableToTypeOf(AmqpConfig{}))
				Expect(rhd.HTTPConf.URL.String()).To(Equal(rhd1HTTP))
				Expect(rhd.AmqpConf.URI.String()).NotTo(Equal(rhd1AMQP))
			})

			It("should have errored", func() {
				Expect(err).To(HaveOccurred())
			})

		})
	})

	Describe("Creating RackHD Instances", func() {
		var err error

		BeforeEach(func() {
			rhd1UUID = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
			rhd1HTTP = "http://10.10.10.10:2020"
			rhd1AMQP = "amqp://localhost/"
		})

		AfterEach(func() {
			db.DeleteTree("rhdman/rhd/" + rhd1UUID)
		})

		JustBeforeEach(func() {
			rhd, err = NewRhd(rhd1UUID, rhd1HTTP, rhd1AMQP)
			if err != nil {
				Fail("Could not create object")
			}
			err = CreateRhd(rhd)
		})

		Context("when valid data is provided", func() {
			It("should create an instance without error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should populate the DB with the correct values", func() {
				pair, err := db.Get("rhdman/rhd/" + rhd1UUID + "/" + "httpConf")
				if err != nil {
					Fail("Unable to get value from the DB for testing")
				}
				Expect(string(pair.Value)).To(Equal(rhd1HTTP))
				pair, err = db.Get("rhdman/rhd/" + rhd1UUID + "/" + "amqpConf")
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
			rhd1, _ := NewRhd(rhd1UUID, rhd1HTTP, rhd1AMQP)
			_ = CreateRhd(rhd1)
			rhd2UUID = "2e650685-bbe7-44f6-9c66-bd466d1bc1ab"
			rhd2HTTP = "http://20.20.20.20:4020"
			rhd2AMQP = "amqp://localhost/"
			rhd2, _ := NewRhd(rhd2UUID, rhd2HTTP, rhd2AMQP)
			_ = CreateRhd(rhd2)
			rhd3UUID = "6c04c204-5c5c-4daf-9d68-7e3f2c76de2d"
			rhd3HTTP = "http://30.30.30.30:6020"
			rhd3AMQP = "amqp://localhost/"
			rhd3, _ := NewRhd(rhd3UUID, rhd3HTTP, rhd3AMQP)
			_ = CreateRhd(rhd3)

		})

		AfterEach(func() {
			db.DeleteTree("rhdman/rhd/" + rhd1UUID)
			db.DeleteTree("rhdman/rhd/" + rhd2UUID)
			db.DeleteTree("rhdman/rhd/" + rhd3UUID)
		})

		Context("Getting All RackHD Instances", func() {

			It("should not have errored", func() {
				_, err := GetAllRhd()
				Expect(err).ToNot(HaveOccurred())
			})

			It("should get all instances in the DB", func() {
				rhds, _ := GetAllRhd()
				Expect(len(rhds)).To(Equal(3))
			})
		})

		Context("Getting RackHD by ID", func() {

			It("should not have errored", func() {
				_, err := GetRhdByID(rhd2UUID)
				Expect(err).ToNot(HaveOccurred())
			})

			It("should return the right RHD", func() {
				rhd, _ := GetRhdByID(rhd2UUID)
				Expect(rhd.ID).To(Equal(rhd2UUID))
				Expect(rhd.HTTPConf.URL.String()).To(Equal(rhd2HTTP))
				Expect(rhd.AmqpConf.URI.String()).To(Equal(rhd2AMQP))
			})
		})

		Context("Getting multiple RackHDs by multiple IDs", func() {

			It("should not have errored", func() {
				arr := []string{rhd2UUID, rhd3UUID}
				_, err := GetRhdsByIDs(arr)
				Expect(err).ToNot(HaveOccurred())
			})

			It("should return only 2 RHDs", func() {
				arr := []string{rhd2UUID, rhd3UUID}
				rhds, _ := GetRhdsByIDs(arr)
				Expect(len(rhds)).To(Equal(2))
			})
		})
	})

	Describe("Deleting RackHD Instances", func() {
		var err error

		BeforeEach(func() {
			rhd1UUID = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
			rhd1HTTP = "http://10.10.10.10:2020"
			rhd1AMQP = "amqp://localhost/"
			rhd1, _ := NewRhd(rhd1UUID, rhd1HTTP, rhd1AMQP)
			_ = CreateRhd(rhd1)
			rhd2UUID = "2e650685-bbe7-44f6-9c66-bd466d1bc1ab"
			rhd2HTTP = "http://20.20.20.20:4020"
			rhd2AMQP = "amqp://localhost/"
			rhd2, _ := NewRhd(rhd2UUID, rhd2HTTP, rhd2AMQP)
			_ = CreateRhd(rhd2)
			rhd3UUID = "6c04c204-5c5c-4daf-9d68-7e3f2c76de2d"
			rhd3HTTP = "http://30.30.30.30:6020"
			rhd3AMQP = "amqp://localhost/"
			rhd3, _ := NewRhd(rhd3UUID, rhd3HTTP, rhd3AMQP)
			_ = CreateRhd(rhd3)
		})

		AfterEach(func() {
			db.DeleteTree("rhdman/rhd/" + rhd1UUID)
			db.DeleteTree("rhdman/rhd/" + rhd2UUID)
			db.DeleteTree("rhdman/rhd/" + rhd3UUID)
		})

		Context("Deleting RackHD Instance by ID", func() {
			It("should not have errored", func() {
				err = DeleteRhdByID(rhd1UUID)
				Expect(err).ToNot(HaveOccurred())
			})

			It("should have removed the instance", func() {
				_ = DeleteRhdByID(rhd1UUID)
				rhds, _ := GetAllRhd()
				Expect(len(rhds)).To(Equal(2))
			})
		})

		Context("Deleting multiple RackHDs by multiple IDs", func() {
			It("should not have errored", func() {
				arr := []string{rhd2UUID, rhd3UUID}
				err := DeleteRhdsByIDs(arr)
				Expect(err).ToNot(HaveOccurred())
			})

			It("should have removed multiple instances", func() {
				arr := []string{rhd2UUID, rhd3UUID}
				_ = DeleteRhdsByIDs(arr)
				rhds, _ := GetAllRhd()
				Expect(len(rhds)).To(Equal(1))
			})
		})
	})
})
