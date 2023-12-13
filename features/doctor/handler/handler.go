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
			c.Logger().Error("Handler : Search Doctor by Name Error : ", err.Error())
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

		result, err := mdl.svc.GetDoctors()
		if err != nil {
			c.Logger().Error("Handler : Get All Process Error : ", err.Error())
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
			c.Logger().Error("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", "Invalid ID"))
		}

		result, err := mdl.svc.GetDoctor(id)

		if err != nil {
			c.Logger().Error("Handler : Get By ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		if result.ID == 0 {
			return c.JSON(http.StatusOK, helper.FormatResponse("Success", "Data is Empty"))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (mdl *DoctorHandler) GetDoctorByUserId() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Error("Handler : Param User ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", "Invalid User ID"))
		}

		result, err := mdl.svc.GetDoctorByUserId(id)

		if err != nil {
			c.Logger().Error("Handler : Get By User ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail", nil))
		}

		if result.ID == 0 {
			return c.JSON(http.StatusOK, helper.FormatResponse("Success", "Data is Empty"))
		}

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
			c.Logger().Error("Handler : Bind Input Error : ", err.Error())
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
		serviceInput.DoctorDescription = input.DoctorDescription
		serviceInput.DoctorNIK = input.DoctorNIK
		serviceInput.DoctorDOB = input.DoctorDOB
		serviceInput.DoctorProvinsi = input.DoctorProvinsi
		serviceInput.DoctorKota = input.DoctorKota
		serviceInput.DoctorNumberPhone = input.DoctorNumberPhone
		serviceInput.DoctorGender = input.DoctorGender

		serviceInput.DoctorAvatar = uploadUrlPhoto
		serviceInput.DoctorMeetLink = input.DoctorMeetLink
		serviceInput.DoctorSIPP = input.DoctorSIPP
		serviceInput.DoctorSTR = input.DoctorSTR
		serviceInput.DoctorSIPPFile = uploadUrlSIPP
		serviceInput.DoctorSTRFile = uploadUrlSTR
		serviceInput.DoctorCV = uploadUrlCV
		serviceInput.DoctorIjazah = uploadUrlIjazah

		serviceInput.DoctorBalance = 0
		serviceInput.DoctorStatus = "request"
		//INPUT REQUEST

		result, err := mdl.svc.CreateDoctor(*serviceInput)

		if err != nil {
			c.Logger().Error("Handler: Input Process Error (CreateDoctor): ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var serviceInputExpertise = new(doctor.DoctorExpertiseRelation)
		serviceInputExpertise.DoctorID = result.ID
		serviceInputExpertise.ExpertiseID = input.DoctorExpertiseID

		resultExpertise, err := mdl.svc.CreateDoctorExpertise(*serviceInputExpertise)

		if err != nil {
			c.Logger().Error("Handler: Input Process Error (CreateDoctorExpertise): ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		if len(input.DoctorWorkdayID) != len(input.DoctorWorkStartTime) || len(input.DoctorWorkdayID) != len(input.DoctorWorkEndTime) {
			c.Logger().Error("Handler: workday, start time, and end time must have the same array length!")
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail: Workday, Start and End time mismatch array length.", nil))
		}

		// Validate array lengths for DoctorEducation
		if len(input.DoctorUniversity) != len(input.DoctorStudyProgram) || len(input.DoctorUniversity) != len(input.DoctorGraduateYear) {
			c.Logger().Error("Handler: university, study program, and graduate year must have the same array length!")
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail: University, Study Program, and Graduate Year mismatch array length.", nil))
		}

		// Validate array lengths for DoctorExperience
		if len(input.DoctorCompany) != len(input.DoctorTitle) ||
			len(input.DoctorCompany) != len(input.DoctorStartDate) || len(input.DoctorCompany) != len(input.DoctorEndDate) {
			c.Logger().Error("Handler: company, title, experience, start date, end date, and is now must have the same array length!")
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail: Company, Title, Experience, Start Date, End Date, and Is Now mismatch array length.", nil))
		}

		// Create DoctorWorkadays objects
		var resultWorkadaysSlice []*doctor.DoctorWorkdays
		for i, workdayID := range input.DoctorWorkdayID {
			// Extract values for the current iteration
			startTime := input.DoctorWorkStartTime[i]
			endTime := input.DoctorWorkEndTime[i]

			// Create a DoctorWorkadays object
			serviceInputWorkadays := doctor.DoctorWorkdays{
				DoctorID:  result.ID,
				WorkdayID: workdayID,
				StartTime: startTime,
				EndTime:   endTime,
			}

			// Call the service to create DoctorWorkadays
			resultWorkadays, err := mdl.svc.CreateDoctorWorkadays(serviceInputWorkadays)
			resultWorkadaysSlice = append(resultWorkadaysSlice, resultWorkadays)

			if err != nil {
				c.Logger().Error("Handler: Input Process Error (CreateDoctorWorkadays): ", err.Error())
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
			}
		}

		// Create DoctorEducation objects
		var resultEducationSlice []*doctor.DoctorEducation

		for i, university := range input.DoctorUniversity {
			// Extract values for the current iteration
			studyProgram := input.DoctorStudyProgram[i]
			graduateYear := input.DoctorGraduateYear[i]
			enrollyear := input.DoctorEnrollYear[i]

			// Create a DoctorEducation object
			serviceInputEducation := doctor.DoctorEducation{
				DoctorID:           result.ID,
				DoctorUniversity:   university,
				DoctorStudyProgram: studyProgram,
				DoctorEnrollYear:   enrollyear,
				DoctorGraduateYear: graduateYear,
			}

			// Call the service to create DoctorEducation
			resultEducation, err := mdl.svc.CreateDoctorEducation(serviceInputEducation)
			resultEducationSlice = append(resultEducationSlice, resultEducation)

			if err != nil {
				c.Logger().Error("Handler: Input Process Error (CreateDoctorEducation): ", err.Error())
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
			}
		}

		// Create DoctorExperience objects
		var resultExperienceSlice []*doctor.DoctorExperience
		for i, company := range input.DoctorCompany {
			// Extract values for the current iteration
			title := input.DoctorTitle[i]
			companyaddress := input.DoctorCompanyAddress[i]
			startDate := input.DoctorStartDate[i]
			endDate := input.DoctorEndDate[i]

			// Create a DoctorExperience object
			serviceInputExperience := doctor.DoctorExperience{
				DoctorID:             result.ID,
				DoctorCompany:        company,
				DoctorTitle:          title,
				DoctorCompanyAddress: companyaddress,
				DoctorStartDate:      startDate,
				DoctorEndDate:        endDate,
			}

			// Call the service to create DoctorExperience
			resultExperience, err := mdl.svc.CreateDoctorExperience(serviceInputExperience)
			resultExperienceSlice = append(resultExperienceSlice, resultExperience)

			if err != nil {
				c.Logger().Error("Handler: Input Process Error (CreateDoctorExperience): ", err.Error())
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
			}
		}

		var response = new(DoctorResponse)

		response.UserID = result.UserID
		response.DoctorName = result.DoctorName
		response.DoctorDescription = result.DoctorDescription
		response.DoctorAvatar = result.DoctorAvatar
		response.DoctorExpertise = resultExpertise.ExpertiseID

		response.DoctorNIK = result.DoctorNIK
		response.DoctorDOB = result.DoctorDOB
		response.DoctorProvinsi = result.DoctorProvinsi
		response.DoctorKota = result.DoctorKota
		response.DoctorNumberPhone = result.DoctorNumberPhone
		response.DoctorGender = result.DoctorGender

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

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success", response))
	}
}

func (mdl *DoctorHandler) UpdateDoctorDatapokok() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Error("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", "Invalid ID"))
		}

		role := mdl.jwt.CheckRole(c)

		fmt.Println(role)
		if role != "Doctor" && role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Fail, you don't have access.", nil))
		}

		var input = new(DoctorRequestDatapokok)
		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler: Bind Input Error: ", err.Error())
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

		var serviceInput = new(doctor.DoctorDatapokokUpdate)

		serviceInput.DoctorName = input.DoctorName
		serviceInput.DoctorDescription = input.DoctorDescription
		serviceInput.DoctorNIK = input.DoctorNIK
		serviceInput.DoctorDOB = input.DoctorDOB
		serviceInput.DoctorProvinsi = input.DoctorProvinsi
		serviceInput.DoctorKota = input.DoctorKota
		serviceInput.DoctorNumberPhone = input.DoctorNumberPhone
		serviceInput.DoctorGender = input.DoctorGender

		serviceInput.DoctorAvatar = uploadUrlPhoto
		serviceInput.DoctorMeetLink = input.DoctorMeetLink
		serviceInput.DoctorSIPP = input.DoctorSIPP
		serviceInput.DoctorSTR = input.DoctorSTR
		serviceInput.DoctorSIPPFile = uploadUrlSIPP
		serviceInput.DoctorSTRFile = uploadUrlSTR
		serviceInput.DoctorCV = uploadUrlCV
		serviceInput.DoctorIjazah = uploadUrlIjazah
		serviceInput.DoctorBalance = input.DoctorBalance
		serviceInput.DoctorStatus = input.DoctorStatus

		serviceInput.DoctorExpertiseID = input.DoctorExpertiseID

		result, err := mdl.svc.UpdateDoctorDatapokok(id, *serviceInput)

		if err != nil {
			c.Logger().Error("Handler: Update Process Error (UpdateDoctor): ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (mdl *DoctorHandler) UpdateDoctorExperience() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Error("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", "Invalid ID"))
		}

		role := mdl.jwt.CheckRole(c)
		fmt.Println(role)

		if role != "Doctor" && role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Fail, you don't have access.", nil))
		}

		var input = new([]DoctorExperience)

		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler: Bind Input Error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var result bool
		var resultUpdate []UpdateResponse

		for _, experience := range *input {
			experienceServiceInput := &doctor.DoctorExperience{
				ID:                   experience.ID,
				DoctorCompany:        experience.DoctorCompany,
				DoctorTitle:          experience.DoctorTitle,
				DoctorCompanyAddress: experience.DoctorCompanyAddress,
				DoctorStartDate:      experience.DoctorStartDate,
				DoctorEndDate:        experience.DoctorEndDate,
			}

			// Update the experience data
			result, err = mdl.svc.UpdateDoctorExperience(int(experience.ID), id, *experienceServiceInput)
			if result == true {
				resultSlice := UpdateResponse{
					ID:          experience.ID,
					Status:      result,
					Description: "Succesfully updated",
				}
				resultUpdate = append(resultUpdate, resultSlice)
			} else {
				resultSlice := UpdateResponse{
					ID:          experience.ID,
					Status:      result,
					Description: "Insufficient permission or no data",
				}
				resultUpdate = append(resultUpdate, resultSlice)
			}

			if err != nil {
				c.Logger().Error("Handler: Update Process Error (UpdateDoctorExperience): ", err.Error())
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
			}
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", resultUpdate))
	}
}

func (mdl *DoctorHandler) UpdateDoctorEducation() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Error("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", "Invalid ID"))
		}

		role := mdl.jwt.CheckRole(c)
		fmt.Println(role)

		if role == "Patient" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Fail, you don't have access.", nil))
		}

		var input = new([]DoctorEducation)
		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler: Bind Input Error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var result bool
		var resultUpdate []UpdateResponse

		for _, education := range *input {
			educationServiceInput := &doctor.DoctorEducation{
				ID:                 education.ID,
				DoctorUniversity:   education.DoctorUniversity,
				DoctorStudyProgram: education.DoctorStudyProgram,
				DoctorEnrollYear:   education.DoctorEnrollYear,
				DoctorGraduateYear: education.DoctorGraduateYear,
			}

			// Update the education data
			result, err = mdl.svc.UpdateDoctorEducation(int(education.ID), id, *educationServiceInput)
			if result == true {
				resultSlice := UpdateResponse{
					ID:          education.ID,
					Status:      result,
					Description: "Succesfully updated",
				}
				resultUpdate = append(resultUpdate, resultSlice)
			} else {
				resultSlice := UpdateResponse{
					ID:          education.ID,
					Status:      result,
					Description: "Insufficient permission or no data",
				}
				resultUpdate = append(resultUpdate, resultSlice)
			}

			if err != nil {
				c.Logger().Error("Handler: Update Process Error (UpdateDoctorEducation): ", err.Error())
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
			}
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", resultUpdate))
	}
}

func (mdl *DoctorHandler) UpdateDoctorWorkdays() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Error("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", "Invalid ID"))
		}

		role := mdl.jwt.CheckRole(c)
		fmt.Println(role)

		if role == "Patient" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Fail, you don't have access.", nil))
		}

		var input = new([]DoctorWorkdays)
		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler: Bind Input Error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var result bool
		var resultUpdate []UpdateResponse

		for _, workday := range *input {

			serviceInput := &doctor.DoctorWorkdays{
				ID:        workday.ID,
				WorkdayID: workday.WorkdayID,
				StartTime: workday.StartTime,
				EndTime:   workday.EndTime,
			}

			// Update the workday data
			result, err = mdl.svc.UpdateDoctorWorkdays(int(workday.ID), id, *serviceInput)
			if result == true {
				resultSlice := UpdateResponse{
					ID:          workday.ID,
					Status:      result,
					Description: "Succesfully updated",
				}
				resultUpdate = append(resultUpdate, resultSlice)
			} else {
				resultSlice := UpdateResponse{
					ID:          workday.ID,
					Status:      result,
					Description: "Insufficient permission or no data",
				}
				resultUpdate = append(resultUpdate, resultSlice)
			}

			if err != nil {
				c.Logger().Error("Handler: Update Process Error (UpdateDoctor): ", err.Error())
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
			}
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", resultUpdate))
	}
}

func (mdl *DoctorHandler) UpdateDoctorRating() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Error("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", "Invalid ID"))
		}

		var paramIDPatient = c.Param("patient")
		patientID, err := strconv.Atoi(paramIDPatient)
		if err != nil {
			c.Logger().Error("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", "Invalid Patient ID"))
		}

		role := mdl.jwt.CheckRole(c)
		fmt.Println(role)

		if role != "Patient" && role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Fail, you don't have access.", nil))
		}

		var input = new(DoctorRating)
		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler: Bind Input Error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		var serviceInput = new(doctor.DoctorRating)

		serviceInput.DoctorReview = input.DoctorReview
		serviceInput.DoctorStarRating = input.DoctorStarRating

		result, err := mdl.svc.UpdateDoctorRating(id, patientID, *serviceInput)

		if err != nil {
			c.Logger().Error("Handler: Update Process Error (UpdateDoctor): ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (mdl *DoctorHandler) DeleteDoctor() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Error("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", "Invalid ID"))
		}

		role := mdl.jwt.CheckRole(c)
		fmt.Println(role)

		if role != "Doctor" && role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Fail, you don't have access.", nil))
		}

		result, err := mdl.svc.DeleteDoctor(id)

		if err != nil {
			c.Logger().Info("Handler: Delete Process Error (Delete Doctor Workdays): ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to create doctor workadays data", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", result))
	}
}

func (mdl *DoctorHandler) DeleteDoctorData() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramType = c.Param("type")
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Error("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", "Invalid ID"))
		}

		role := mdl.jwt.CheckRole(c)
		fmt.Println(role)

		if role != "Doctor" && role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Fail, you don't have access.", nil))
		}

		if paramType == "workday" {

			result, err := mdl.svc.DeleteDoctorWorkdays(id)

			if err != nil {
				c.Logger().Info("Handler: Delete Process Error (Delete Doctor Workdays): ", err.Error())
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to create doctor workadays data", nil))
			}

			return c.JSON(http.StatusCreated, helper.FormatResponse("Success", result))

		} else if paramType == "education" {

			result, err := mdl.svc.DeleteDoctorEducation(id)

			if err != nil {
				c.Logger().Info("Handler: Delete Process Error (Delete Doctor Education): ", err.Error())
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to create doctor education data", nil))
			}

			return c.JSON(http.StatusCreated, helper.FormatResponse("Success", result))

		} else if paramType == "experience" {

			result, err := mdl.svc.DeleteDoctorExperience(id)

			if err != nil {
				c.Logger().Info("Handler: Delete Process Error (Delete Doctor Experience): ", err.Error())
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to create doctor experience data", nil))
			}

			return c.JSON(http.StatusCreated, helper.FormatResponse("Success", result))

		} else if paramType == "rating" {

			result, err := mdl.svc.DeleteDoctorRating(id)

			if err != nil {
				c.Logger().Info("Handler: Delete Process Error (Delete Doctor Rating): ", err.Error())
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to create doctor experience data", nil))
			}

			return c.JSON(http.StatusCreated, helper.FormatResponse("Success", result))

		}

		return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail, type or id not found", nil))
	}
}

func (mdl *DoctorHandler) InsertDataDoctor() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramType = c.Param("type")

		role := mdl.jwt.CheckRole(c)
		fmt.Println(role)

		if role != "Doctor" && role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Fail, you don't have access.", nil))
		}

		if paramType == "workday" {

			var input = new(DoctorWorkdays)
			if err := c.Bind(input); err != nil {
				c.Logger().Info("Handler: Bind Input Error: ", err.Error())
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
			}

			var serviceInput = new(doctor.DoctorWorkdays)
			serviceInput.DoctorID = input.DoctorID
			serviceInput.WorkdayID = input.WorkdayID
			serviceInput.StartTime = input.StartTime
			serviceInput.EndTime = input.EndTime

			result, err := mdl.svc.CreateDoctorWorkadays(*serviceInput)

			if err != nil {
				c.Logger().Info("Handler: Insert Process Error (Insert Doctor Workdays): ", err.Error())
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to create doctor workadays data", nil))
			}

			return c.JSON(http.StatusCreated, helper.FormatResponse("Success", result))

		} else if paramType == "education" {

			var input = new(DoctorEducation)
			if err := c.Bind(input); err != nil {
				c.Logger().Info("Handler: Bind Input Error: ", err.Error())
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
			}

			var serviceInput = new(doctor.DoctorEducation)
			serviceInput.DoctorID = input.DoctorID
			serviceInput.DoctorStudyProgram = input.DoctorStudyProgram
			serviceInput.DoctorUniversity = input.DoctorUniversity
			serviceInput.DoctorEnrollYear = input.DoctorEnrollYear
			serviceInput.DoctorGraduateYear = input.DoctorGraduateYear

			result, err := mdl.svc.CreateDoctorEducation(*serviceInput)

			if err != nil {
				c.Logger().Info("Handler: Insert Process Error (Insert Doctor Education): ", err.Error())
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to create doctor education data", nil))
			}

			return c.JSON(http.StatusCreated, helper.FormatResponse("Success", result))

		} else if paramType == "experience" {
			var input = new(DoctorExperience)
			if err := c.Bind(input); err != nil {
				c.Logger().Info("Handler: Bind Input Error: ", err.Error())
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
			}

			var serviceInput = new(doctor.DoctorExperience)
			serviceInput.DoctorID = input.DoctorID
			serviceInput.DoctorCompanyAddress = input.DoctorCompanyAddress
			serviceInput.DoctorCompany = input.DoctorCompany
			serviceInput.DoctorTitle = input.DoctorTitle
			serviceInput.DoctorStartDate = input.DoctorStartDate
			serviceInput.DoctorEndDate = input.DoctorEndDate

			result, err := mdl.svc.CreateDoctorExperience(*serviceInput)

			if err != nil {
				c.Logger().Info("Handler: Insert Process Error (Insert Doctor Experience): ", err.Error())
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to create doctor experience data", nil))
			}

			return c.JSON(http.StatusCreated, helper.FormatResponse("Success", result))

		}

		return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail, insert type not found", nil))
	}
}
