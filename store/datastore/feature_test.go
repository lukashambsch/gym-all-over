package datastore_test

import (
	"github.com/lukashambsch/anygym.api/models"
	"github.com/lukashambsch/anygym.api/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Feature db interactions", func() {
	var featureID int64 = 1

	Describe("GetFeatureList", func() {
		var features []models.Feature

		Describe("Successful call", func() {
			BeforeEach(func() {
				features, _ = datastore.GetFeatureList("")
			})

			It("should return a list of features", func() {
				Expect(len(features)).To(Equal(24))
			})
		})
	})

	Describe("GetFeature", func() {
		var feature *models.Feature

		Describe("Successful call", func() {
			It("should return the correct feature", func() {
				feature, _ = datastore.GetFeature(featureID)
				Expect(feature.FeatureID).To(Equal(featureID))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentID int64 = 5000
				err           error
			)

			BeforeEach(func() {
				feature, err = datastore.GetFeature(nonExistentID)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil plan", func() {
				Expect(feature).To(BeNil())
			})
		})
	})

	Describe("GetFeatureCount", func() {
		var count *int

		Describe("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetFeatureCount("")
			})

			It("should return the correct count", func() {
				Expect(*count).To(Equal(24))
			})
		})
	})

	Describe("CreateFeature", func() {
		var (
			featureName string = "Test Name"
			feature     models.Feature
			created     *models.Feature
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				feature = models.Feature{FeatureName: featureName, FeatureDescription: "Test Description"}
				created, _ = datastore.CreateFeature(feature)
			})

			AfterEach(func() {
				datastore.DeleteFeature(created.FeatureID)
			})

			It("should return the created feature", func() {
				Expect(created.FeatureName).To(Equal(featureName))
			})

			It("should add a feature to the db", func() {
				newFeature, _ := datastore.GetFeature(created.FeatureID)
				Expect(newFeature.FeatureName).To(Equal(featureName))
			})
		})

		Describe("Unsuccessful call", func() {
			var created *models.Feature

			AfterEach(func() {
				datastore.DeleteFeature(created.FeatureID)
			})

			It("should return an error object if feature is not unique", func() {
				usrRole := models.Feature{FeatureName: featureName, FeatureDescription: "Test Description"}
				created, _ = datastore.CreateFeature(usrRole)
				_, err := datastore.CreateFeature(usrRole)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("UpdateFeature", func() {
		var (
			featureName string = "New Name"
			feature     models.Feature
			created     *models.Feature
			updated     *models.Feature
			err         error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				feature = models.Feature{FeatureName: featureName, FeatureDescription: "Test Description"}
				created, _ = datastore.CreateFeature(models.Feature{FeatureName: "Test"})
				updated, _ = datastore.UpdateFeature(created.FeatureID, feature)
			})

			AfterEach(func() {
				datastore.DeleteFeature(updated.FeatureID)
			})

			It("should return the updated feature", func() {
				Expect(updated.FeatureName).To(Equal(featureName))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				feature = models.Feature{}
				updated, err = datastore.UpdateFeature(10000, feature)
			})

			It("should return an error object if feature to update doesn't exist", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil feature", func() {
				Expect(updated).To(BeNil())
			})
		})
	})

	Describe("DeleteFeature", func() {
		Describe("Successful call", func() {
			It("should return nil", func() {
				created, _ := datastore.CreateFeature(models.Feature{FeatureName: "Test"})
				err := datastore.DeleteFeature(created.FeatureID)
				Expect(err).To(BeNil())
			})
		})
	})
})
