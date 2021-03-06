package datastore_test

import (
	"github.com/lukashambsch/anygym.api/models"
	"github.com/lukashambsch/anygym.api/store/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Plan db interactions", func() {
	var one, two, three, four *models.Plan

	BeforeEach(func() {
		one, _ = datastore.CreatePlan(models.Plan{PlanName: "Testing"})
		two, _ = datastore.CreatePlan(models.Plan{PlanName: "Testing Two"})
		three, _ = datastore.CreatePlan(models.Plan{PlanName: "Testing Three"})
		four, _ = datastore.CreatePlan(models.Plan{PlanName: "Testing Four"})
	})

	AfterEach(func() {
		datastore.DeletePlan(one.PlanID)
		datastore.DeletePlan(two.PlanID)
		datastore.DeletePlan(three.PlanID)
		datastore.DeletePlan(four.PlanID)
	})

	Describe("GetPlanList", func() {
		var plans []models.Plan

		Describe("Successful call", func() {
			BeforeEach(func() {
				plans, _ = datastore.GetPlanList("")
			})

			It("should return a list of plans", func() {
				Expect(len(plans)).To(Equal(6))
			})
		})
	})

	Describe("GetPlan", func() {
		var plan *models.Plan

		Describe("Successful call", func() {
			It("should return the correct plan", func() {
				plan, _ = datastore.GetPlan(one.PlanID)
				Expect(plan.PlanID).To(Equal(one.PlanID))
			})
		})

		Describe("Unsuccessful call", func() {
			var (
				nonExistentID int64 = 5
				err           error
			)

			BeforeEach(func() {
				plan, err = datastore.GetPlan(nonExistentID)
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil plan", func() {
				Expect(plan).To(BeNil())
			})
		})
	})

	Describe("GetPlanCount", func() {
		var count *int

		Describe("Successful call", func() {
			BeforeEach(func() {
				count, _ = datastore.GetPlanCount("")
			})

			It("should return the correct count", func() {
				Expect(*count).To(Equal(6))
			})
		})
	})

	Describe("CreatePlan", func() {
		var (
			planName string = "New Plan"
			plan     models.Plan
			created  *models.Plan
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				plan = models.Plan{PlanName: planName}
				created, _ = datastore.CreatePlan(plan)
			})

			AfterEach(func() {
				datastore.DeletePlan(created.PlanID)
			})

			It("should return the created plan", func() {
				Expect(created.PlanName).To(Equal(planName))
			})

			It("should add a plan to the db", func() {
				newPlan, _ := datastore.GetPlan(created.PlanID)
				Expect(newPlan.PlanName).To(Equal(planName))
			})
		})

		Describe("Unsuccessful call", func() {
			var created *models.Plan

			AfterEach(func() {
				datastore.DeletePlan(created.PlanID)
			})

			It("should return an error object if plan is not unique", func() {
				name := "Test Name"
				pln := models.Plan{PlanName: name}
				created, _ = datastore.CreatePlan(pln)
				_, err := datastore.CreatePlan(pln)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("UpdatePlan", func() {
		var (
			planName string = "Anytime"
			plan     models.Plan
			created  *models.Plan
			updated  *models.Plan
			err      error
		)

		Describe("Successful call", func() {
			BeforeEach(func() {
				plan = models.Plan{PlanName: planName}
				created, _ = datastore.CreatePlan(models.Plan{PlanName: "Daily"})
				updated, _ = datastore.UpdatePlan(created.PlanID, plan)
			})

			AfterEach(func() {
				datastore.DeletePlan(updated.PlanID)
			})

			It("should return the updated plan", func() {
				Expect(updated.PlanName).To(Equal(planName))
			})
		})

		Describe("Unsuccessful call", func() {
			BeforeEach(func() {
				plan = models.Plan{PlanName: "Daily"}
				updated, err = datastore.UpdatePlan(10000, plan)
			})

			It("should return an error object if plan to update doesn't exist", func() {
				Expect(err).ToNot(BeNil())
			})

			It("should return a nil plan", func() {
				Expect(updated).To(BeNil())
			})
		})
	})

	Describe("DeletePlan", func() {
		Describe("Successful call", func() {
			It("should return nil", func() {
				err := datastore.DeletePlan(one.PlanID)
				Expect(err).To(BeNil())
			})
		})
	})
})
