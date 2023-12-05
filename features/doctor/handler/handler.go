package handler

import (
	"FinalProject/features/doctor"
	"FinalProject/helper"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DoctorHandler struct {
	svc doctor.DoctorServiceInterface
	jwt helper.JWTInterface
}

func NewHandlerDoctor(service doctor.DoctorServiceInterface, jwt helper.JWTInterface) doctor.DoctorHandlerInterface {
	return &DoctorHandler{
		svc: service,
		jwt: jwt,
	}
}

func (mdl *DoctorHandler) SearchDoctor() echo.HandlerFunc {
	return func(c echo.Context) error {

		name := c.QueryParam("name")
		if name == "" {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed", "Name parameter is required"))
		}

		result, err := mdl.svc.SearchDoctor(name)

		if err != nil {
			c.Logger().Fatal("Handler : Search Doctor by Name Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		if len(result) == 0 {
			return c.JSON(http.StatusOK, helper.FormatResponse("Success", "No matching doctors found"))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (mdl *DoctorHandler) GetDoctors() echo.HandlerFunc {
	return func(c echo.Context) error {
		// role := mdl.jwt.CheckRole(c)
		// fmt.Println(role)
		// if role != "Admin" || role != "Doctor" {
		// 	return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized", nil))
		// }

		result, err := mdl.svc.GetDoctors()

		if err != nil {
			c.Logger().Fatal("Handler : Get All Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		if len(result) == 0 {
			return c.JSON(http.StatusOK, helper.FormatResponse("Success", "Data is Empty"))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (mdl *DoctorHandler) GetDoctor() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Info("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", "Invalid ID"))
		}

		result, err := mdl.svc.GetDoctor(id)
		// resultExperience, err := mdl.svc.GetDoctorExperience(id)
		// resultWorkadays, err := mdl.svc.GetDoctorWorkadays(id)
		// resultEducation, err := mdl.svc.GetDoctorEducation(id)

		if err != nil {
			c.Logger().Fatal("Handler : Get By ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		// mapAllData := map[string]interface{}{
		// 	"doctor":     result,
		// 	"workadays":  resultWorkadays,
		// 	"experience": resultExperience,
		// 	"education":  resultEducation,
		// }

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (mdl *DoctorHandler) CreateDoctor() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mdl.jwt.CheckRole(c)
		getID, err := mdl.jwt.GetID(c)
		fmt.Println(role)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Fail, cant get ID from JWT", nil))
		}

		var input = new(DoctorRequest)

		if err := c.Bind(input); err != nil {
			c.Logger().Fatal("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		formHeaderPhoto, err := c.FormFile("doctor_avatar")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed, Select a File for Upload Avatar", nil))
		}

		formHeaderSIPP, err := c.FormFile("doctor_sipp_file")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed, Select a File for Upload SIPP", nil))
		}

		formHeaderSTR, err := c.FormFile("doctor_str_file")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed, Select a File for Upload STR", nil))
		}

		formHeaderCV, err := c.FormFile("doctor_cv")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed, Select a File for Upload CV", nil))
		}

		formHeaderIjazah, err := c.FormFile("doctor_ijazah")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed, Select a File for Upload Ijazah", nil))
		}

		formPhoto, err := formHeaderPhoto.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed FormHeaderPhoto", nil))
		}

		formSIPP, err := formHeaderSIPP.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed FormHeaderSIPP", nil))
		}

		formSTR, err := formHeaderSTR.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed FormHeaderSTR", nil))
		}

		formCV, err := formHeaderCV.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed FormHeaderCV", nil))
		}

		formIjazah, err := formHeaderIjazah.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed FormHeaderIjazah", nil))
		}

		uploadUrlPhoto, err := mdl.svc.DoctorAvatarUpload(doctor.DoctorAvatarPhoto{DoctorAvatar: formPhoto})
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed Upload Avatar", nil))
		}

		uploadUrlSIPP, err := mdl.svc.DoctorSIPPUpload(doctor.DoctorSIPPFileDataModel{DoctorSIPPFile: formSIPP})
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed Upload SIPP", nil))
		}

		uploadUrlSTR, err := mdl.svc.DoctorSTRUpload(doctor.DoctorSTRFileDataModel{DoctorSTRFile: formSTR})
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed Upload STR", nil))
		}

		uploadUrlCV, err := mdl.svc.DoctorCVUpload(doctor.DoctorCVDataModel{DoctorCV: formCV})
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed Upload CV", nil))
		}

		uploadUrlIjazah, err := mdl.svc.DoctorIjazahUpload(doctor.DoctorIjazahDataModel{DoctorIjazah: formIjazah})
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed Upload Ijazah", nil))
		}

		var serviceInput = new(doctor.Doctor)

		serviceInput.UserID = getID
		serviceInput.DoctorName = input.DoctorName
		serviceInput.DoctorExperienced = input.DoctorExperienced
		serviceInput.DoctorDescription = input.DoctorDescription
		serviceInput.DoctorAvatar = uploadUrlPhoto

		serviceInput.DoctorOfficeName = input.DoctorOfficeName
		serviceInput.DoctorOfficeAddress = input.DoctorOfficeAddress
		serviceInput.DoctorOfficeCity = input.DoctorOfficeCity
		serviceInput.DoctorMeetLink = input.DoctorMeetLink

		serviceInput.DoctorSIPPFile = uploadUrlSIPP
		serviceInput.DoctorSTRFile = uploadUrlSTR
		serviceInput.DoctorCV = uploadUrlCV
		serviceInput.DoctorIjazah = uploadUrlIjazah

		serviceInput.DoctorBalance = 0
		serviceInput.DoctorStatus = "request"
		//INPUT REQUEST

		result, err := mdl.svc.CreateDoctor(*serviceInput)

		if err != nil {
			c.Logger().Info("Handler: Input Process Error (CreateDoctor): ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var serviceInputExpertise = new(doctor.DoctorExpertiseRelation)
		serviceInputExpertise.DoctorID = result.ID                  //...
		serviceInputExpertise.ExpertiseID = input.DoctorExpertiseID //...

		resultExpertise, err := mdl.svc.CreateDoctorExpertise(*serviceInputExpertise)

		if err != nil {
			c.Logger().Info("Handler: Input Process Error (CreateDoctorExpertise): ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		if len(input.DoctorWorkdayID) != len(input.DoctorWorkStartTime) || len(input.DoctorWorkdayID) != len(input.DoctorWorkEndTime) {
			c.Logger().Info("Handler: workday, start time, and end time must have the same array length!")
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail: Workday, Start and End time mismatch array length.", nil))
		}

		// Validate array lengths for DoctorEducation
		if len(input.DoctorUniversity) != len(input.DoctorStudyProgram) || len(input.DoctorUniversity) != len(input.DoctorGraduateYear) {
			c.Logger().Info("Handler: university, study program, and graduate year must have the same array length!")
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail: University, Study Program, and Graduate Year mismatch array length.", nil))
		}

		// Validate array lengths for DoctorExperience
		if len(input.DoctorCompany) != len(input.DoctorTitle) || len(input.DoctorCompany) != len(input.DoctorExperienceDescription) ||
			len(input.DoctorCompany) != len(input.DoctorStartDate) || len(input.DoctorCompany) != len(input.DoctorEndDate) ||
			len(input.DoctorCompany) != len(input.DoctorIsNow) {
			c.Logger().Info("Handler: company, title, experience, start date, end date, and is now must have the same array length!")
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail: Company, Title, Experience, Start Date, End Date, and Is Now mismatch array length.", nil))
		}

		// Create DoctorWorkadays objects
		var resultWorkadaysSlice []*doctor.DoctorWorkadays
		for i, workdayID := range input.DoctorWorkdayID {
			// Extract values for the current iteration
			startTime := input.DoctorWorkStartTime[i]
			endTime := input.DoctorWorkEndTime[i]

			// Create a DoctorWorkadays object
			serviceInputWorkadays := doctor.DoctorWorkadays{
				DoctorID:  result.ID,
				WorkdayID: workdayID,
				StartTime: startTime,
				EndTime:   endTime,
			}

			// Call the service to create DoctorWorkadays
			resultWorkadays, err := mdl.svc.CreateDoctorWorkadays(serviceInputWorkadays)
			resultWorkadaysSlice = append(resultWorkadaysSlice, resultWorkadays)

			if err != nil {
				c.Logger().Info("Handler: Input Process Error (CreateDoctorWorkadays): ", err.Error())
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
			}
		}

		// Create DoctorEducation objects
		var resultEducationSlice []*doctor.DoctorEducation

		for i, university := range input.DoctorUniversity {
			// Extract values for the current iteration
			studyProgram := input.DoctorStudyProgram[i]
			graduateYear := input.DoctorGraduateYear[i]

			// Create a DoctorEducation object
			serviceInputEducation := doctor.DoctorEducation{
				DoctorID:           result.ID,
				DoctorUniversity:   university,
				DoctorStudyProgram: studyProgram,
				DoctorGraduateYear: graduateYear,
			}

			// Call the service to create DoctorEducation
			resultEducation, err := mdl.svc.CreateDoctorEducation(serviceInputEducation)
			resultEducationSlice = append(resultEducationSlice, resultEducation)

			if err != nil {
				c.Logger().Info("Handler: Input Process Error (CreateDoctorEducation): ", err.Error())
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
			}
		}

		// Create DoctorExperience objects
		var resultExperienceSlice []*doctor.DoctorExperience
		for i, company := range input.DoctorCompany {
			// Extract values for the current iteration
			title := input.DoctorTitle[i]
			description := input.DoctorExperienceDescription[i]
			startDate := input.DoctorStartDate[i]
			endDate := input.DoctorEndDate[i]
			isNow := input.DoctorIsNow[i]

			// Create a DoctorExperience object
			serviceInputExperience := doctor.DoctorExperience{
				DoctorID:                    result.ID,
				DoctorCompany:               company,
				DoctorTitle:                 title,
				DoctorExperienceDescription: description,
				DoctorStartDate:             startDate,
				DoctorEndDate:               endDate,
				DoctorIsNow:                 isNow,
			}

			// Call the service to create DoctorExperience
			resultExperience, err := mdl.svc.CreateDoctorExperience(serviceInputExperience)
			resultExperienceSlice = append(resultExperienceSlice, resultExperience)

			if err != nil {
				c.Logger().Info("Handler: Input Process Error (CreateDoctorExperience): ", err.Error())
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
			}
		}

		var response = new(DoctorResponse)

		// response.ID = result.
		response.UserID = result.UserID
		response.DoctorName = result.DoctorName
		response.DoctorExpertise = resultExpertise.ExpertiseID
		// response.DoctorExperience = result.DoctorExperience
		// response.DoctorDescription = result.DoctorDescription
		response.DoctorAvatar = result.DoctorAvatar
		response.DoctorOfficeName = result.DoctorOfficeName
		response.DoctorOfficeAddress = result.DoctorOfficeAddress
		response.DoctorOfficeCity = result.DoctorOfficeCity
		response.DoctorMeetLink = result.DoctorMeetLink
		response.DoctorSIPPFile = result.DoctorSIPPFile
		response.DoctorSTRFile = result.DoctorSTRFile
		response.DoctorCV = result.DoctorCV
		response.DoctorIjazah = result.DoctorIjazah
		response.DoctorBalance = result.DoctorBalance
		response.DoctorStatus = result.DoctorStatus

		response.DoctorWorkday = resultWorkadaysSlice
		response.DoctorEducation = resultEducationSlice
		response.DoctorExperience = resultExperienceSlice

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", response))
	}
}

// func (mdl *DoctorHandler) UpdateDoctor() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		role := mdl.jwt.CheckRole(c)
// 		getID, err := mdl.jwt.GetID(c)
// 		fmt.Println(role)
// 		if err != nil {
// 			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Fail, cant get ID from JWT", nil))
// 		}

// 		var input = new(DoctorRequest)
// 		if err := c.Bind(input); err != nil {
// 			c.Logger().Fatal("Handler: Bind Input Error: ", err.Error())
// 			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
// 		}

// 		formHeaderPhoto, err := c.FormFile("doctor_avatar")
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed, Select a File for Upload Avatar", nil))
// 		}

// 		formHeaderSIPP, err := c.FormFile("doctor_sipp_file")
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed, Select a File for Upload SIPP", nil))
// 		}

// 		formHeaderSTR, err := c.FormFile("doctor_str_file")
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed, Select a File for Upload STR", nil))
// 		}

// 		formHeaderCV, err := c.FormFile("doctor_cv")
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed, Select a File for Upload CV", nil))
// 		}

// 		formHeaderIjazah, err := c.FormFile("doctor_ijazah")
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed, Select a File for Upload Ijazah", nil))
// 		}

// 		formPhoto, err := formHeaderPhoto.Open()
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed FormHeaderPhoto", nil))
// 		}

// 		formSIPP, err := formHeaderSIPP.Open()
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed FormHeaderSIPP", nil))
// 		}

// 		formSTR, err := formHeaderSTR.Open()
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed FormHeaderSTR", nil))
// 		}

// 		formCV, err := formHeaderCV.Open()
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed FormHeaderCV", nil))
// 		}

// 		formIjazah, err := formHeaderIjazah.Open()
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed FormHeaderIjazah", nil))
// 		}

// 		uploadUrlPhoto, err := mdl.svc.DoctorAvatarUpload(doctor.DoctorAvatarPhoto{DoctorAvatar: formPhoto})
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed Upload Avatar", nil))
// 		}

// 		uploadUrlSIPP, err := mdl.svc.DoctorSIPPUpload(doctor.DoctorSIPPFileDataModel{DoctorSIPPFile: formSIPP})
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed Upload SIPP", nil))
// 		}

// 		uploadUrlSTR, err := mdl.svc.DoctorSTRUpload(doctor.DoctorSTRFileDataModel{DoctorSTRFile: formSTR})
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed Upload STR", nil))
// 		}

// 		uploadUrlCV, err := mdl.svc.DoctorCVUpload(doctor.DoctorCVDataModel{DoctorCV: formCV})
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed Upload CV", nil))
// 		}

// 		uploadUrlIjazah, err := mdl.svc.DoctorIjazahUpload(doctor.DoctorIjazahDataModel{DoctorIjazah: formIjazah})
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed Upload Ijazah", nil))
// 		}

// 		var serviceInput = new(doctor.Doctor)

// 		serviceInput.UserID = getID
// 		serviceInput.DoctorName = input.DoctorName
// 		serviceInput.DoctorExperienced = input.DoctorExperienced
// 		serviceInput.DoctorDescription = input.DoctorDescription
// 		serviceInput.DoctorAvatar = uploadUrlPhoto

// 		serviceInput.DoctorOfficeName = input.DoctorOfficeName
// 		serviceInput.DoctorOfficeAddress = input.DoctorOfficeAddress
// 		serviceInput.DoctorOfficeCity = input.DoctorOfficeCity
// 		serviceInput.DoctorMeetLink = input.DoctorMeetLink

// 		serviceInput.DoctorSIPPFile = uploadUrlSIPP
// 		serviceInput.DoctorSTRFile = uploadUrlSTR
// 		serviceInput.DoctorCV = uploadUrlCV
// 		serviceInput.DoctorIjazah = uploadUrlIjazah

// 		result, err := mdl.svc.UpdateDoctorMainData(*serviceInput)

// 		if err != nil {
// 			c.Logger().Info("Handler: Update Process Error (UpdateDoctor): ", err.Error())
// 			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
// 		}

// 		// Update DoctorExpertise
// 		var serviceInputExpertise = new(doctor.DoctorExpertiseRelation)
// 		serviceInputExpertise.DoctorID = result.ID
// 		serviceInputExpertise.ExpertiseID = input.DoctorExpertiseID

// 		_, err = mdl.svc.UpdateDoctorExpertise(*serviceInputExpertise)

// 		if err != nil {
// 			c.Logger().Info("Handler: Input Process Error (UpdateDoctorExpertise): ", err.Error())
// 			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
// 		}

// 		// Your existing code for updating DoctorWorkadays, DoctorEducation, and DoctorExperience

// 		var response = new(DoctorResponse)
// 		response.UserID = result.UserID
// 		response.DoctorName = result.DoctorName
// 		response.DoctorExpertise = input.DoctorExpertiseID

// 		return c.JSON(http.StatusOK, helper.FormatResponse("Success", response))
// 	}
// }
