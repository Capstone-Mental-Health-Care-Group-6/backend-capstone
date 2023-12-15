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

func (mdl *DoctorHandler) GetDoctors() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.QueryParam("name")

		result, err := mdl.svc.GetDoctors(name)
		if err != nil {
			c.Logger().Error("Handler : Error to get doctor : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail, error to get doctors", nil))
		}

		if len(result) == 0 {
			return c.JSON(http.StatusOK, helper.FormatResponse("Success", "Doctors data is Empty"))
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
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid doctor id", nil))
		}

		result, err := mdl.svc.GetDoctor(id)

		if err != nil {
			c.Logger().Error("Handler : Get By ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail to get doctor id", nil))
		}

		if result.ID == 0 {
			return c.JSON(http.StatusOK, helper.FormatResponse("Success", "Doctor data is empty"))
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
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", "Invalid user id"))
		}

		result, err := mdl.svc.GetDoctorByUserId(id)

		if err != nil {
			c.Logger().Error("Handler : Get By User ID Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Fail to get user id", nil))
		}

		if result.ID == 0 {
			return c.JSON(http.StatusOK, helper.FormatResponse("Success", "Doctor data is empty"))
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
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Fail to get id from jwt", nil))
		}

		var input = new(DoctorRequest)

		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler : Bind Input Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail", nil))
		}

		formHeaderPhoto, err := c.FormFile("doctor_avatar")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Required to upload avatar", nil))
		}

		formHeaderSIPP, err := c.FormFile("doctor_sipp_file")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Required to upload SIPP", nil))
		}

		formHeaderSTR, err := c.FormFile("doctor_str_file")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Required to upload STR", nil))
		}

		formHeaderCV, err := c.FormFile("doctor_cv")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Required to upload CV", nil))
		}

		formHeaderIjazah, err := c.FormFile("doctor_ijazah")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Required to upload Ijazah", nil))
		}

		formPhoto, err := formHeaderPhoto.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to open avatar", nil))
		}

		formSIPP, err := formHeaderSIPP.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to open SIPP", nil))
		}

		formSTR, err := formHeaderSTR.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to open STR", nil))
		}

		formCV, err := formHeaderCV.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to open CV", nil))
		}

		formIjazah, err := formHeaderIjazah.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to open Ijazah", nil))
		}

		uploadUrlPhoto, err := mdl.svc.DoctorAvatarUpload(doctor.DoctorAvatarPhoto{DoctorAvatar: formPhoto})
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed to upload Avatar", nil))
		}

		uploadUrlSIPP, err := mdl.svc.DoctorSIPPUpload(doctor.DoctorSIPPFileDataModel{DoctorSIPPFile: formSIPP})
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed to upload SIPP", nil))
		}

		uploadUrlSTR, err := mdl.svc.DoctorSTRUpload(doctor.DoctorSTRFileDataModel{DoctorSTRFile: formSTR})
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed to upload STR", nil))
		}

		uploadUrlCV, err := mdl.svc.DoctorCVUpload(doctor.DoctorCVDataModel{DoctorCV: formCV})
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed to upload CV", nil))
		}

		uploadUrlIjazah, err := mdl.svc.DoctorIjazahUpload(doctor.DoctorIjazahDataModel{DoctorIjazah: formIjazah})
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed to upload Ijazah", nil))
		}

		resultMeetlink, err := mdl.svc.GetMeetLink()

		if err != nil {
			c.Logger().Error("Handler: Get meet link error (Get meet link): ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to create meet link", nil))
		}

		var serviceInput = new(doctor.Doctor)

		if resultMeetlink == "" {
			serviceInput.DoctorMeetLink = "https://meet.google.com/(Create by yourself)"
		} else {
			serviceInput.DoctorMeetLink = resultMeetlink
		}

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

		serviceInput.DoctorSIPP = input.DoctorSIPP
		serviceInput.DoctorSTR = input.DoctorSTR
		serviceInput.DoctorSIPPFile = uploadUrlSIPP
		serviceInput.DoctorSTRFile = uploadUrlSTR
		serviceInput.DoctorCV = uploadUrlCV
		serviceInput.DoctorIjazah = uploadUrlIjazah

		serviceInput.DoctorBalance = 0
		serviceInput.DoctorStatus = "Request"
		//INPUT REQUEST

		result, err := mdl.svc.CreateDoctor(*serviceInput)

		if err != nil {
			c.Logger().Error("Handler: Input Process Error (CreateDoctor): ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to create data doctor", nil))
		}

		var serviceInputExpertise = new(doctor.DoctorExpertiseRelation)
		serviceInputExpertise.DoctorID = result.ID
		serviceInputExpertise.ExpertiseID = input.DoctorExpertiseID

		resultExpertise, err := mdl.svc.CreateDoctorExpertise(*serviceInputExpertise)

		if err != nil {
			c.Logger().Error("Handler: Input Process Error (CreateDoctorExpertise): ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to create data doctor expertise", nil))
		}

		if len(input.DoctorWorkdayID) != len(input.DoctorWorkStartTime) || len(input.DoctorWorkdayID) != len(input.DoctorWorkEndTime) {
			c.Logger().Error("Handler: workday, start time, and end time must have the same array length!")
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("workday, start time, and end time must have the same array length", nil))
		}

		// Validate array lengths for DoctorEducation
		if len(input.DoctorUniversity) != len(input.DoctorStudyProgram) || len(input.DoctorUniversity) != len(input.DoctorGraduateYear) {
			c.Logger().Error("Handler: university, study program, and graduate year must have the same array length!")
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("university, study program, and graduate year must have the same array length", nil))
		}

		// Validate array lengths for DoctorExperience
		if len(input.DoctorCompany) != len(input.DoctorTitle) ||
			len(input.DoctorCompany) != len(input.DoctorStartDate) || len(input.DoctorCompany) != len(input.DoctorEndDate) {
			c.Logger().Error("Handler: company, title, experience, start date, end date, and is now must have the same array length!")
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("company, title, experience, start date, end date, and is now must have the same array length", nil))
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
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to create data doctor workdays", nil))
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
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to create data doctor education", nil))
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
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to create data doctor experience", nil))
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
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to get doctor id", nil))
		}

		role := mdl.jwt.CheckRole(c)

		fmt.Println(role)
		if role != "Doctor" && role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Insufficient permission", nil))
		}

		var input = new(DoctorRequestDatapokok)
		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler: Bind Input Error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to bind request", nil))
		}

		formHeaderPhoto, err := c.FormFile("doctor_avatar")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Required to upload Avatar", nil))
		}

		formHeaderSIPP, err := c.FormFile("doctor_sipp_file")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Required to upload SIPP", nil))
		}

		formHeaderSTR, err := c.FormFile("doctor_str_file")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Required to upload STR", nil))
		}

		formHeaderCV, err := c.FormFile("doctor_cv")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Required to upload CV", nil))
		}

		formHeaderIjazah, err := c.FormFile("doctor_ijazah")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Required to upload Ijazah", nil))
		}

		formPhoto, err := formHeaderPhoto.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to open Photo", nil))
		}

		formSIPP, err := formHeaderSIPP.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to open SIPP", nil))
		}

		formSTR, err := formHeaderSTR.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to open STR", nil))
		}

		formCV, err := formHeaderCV.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to open CV", nil))
		}

		formIjazah, err := formHeaderIjazah.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Failed to open Ijazah", nil))
		}

		uploadUrlPhoto, err := mdl.svc.DoctorAvatarUpload(doctor.DoctorAvatarPhoto{DoctorAvatar: formPhoto})
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed to upload Avatar", nil))
		}

		uploadUrlSIPP, err := mdl.svc.DoctorSIPPUpload(doctor.DoctorSIPPFileDataModel{DoctorSIPPFile: formSIPP})
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed to upload SIPP", nil))
		}

		uploadUrlSTR, err := mdl.svc.DoctorSTRUpload(doctor.DoctorSTRFileDataModel{DoctorSTRFile: formSTR})
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed to upload STR", nil))
		}

		uploadUrlCV, err := mdl.svc.DoctorCVUpload(doctor.DoctorCVDataModel{DoctorCV: formCV})
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed to upload CV", nil))
		}

		uploadUrlIjazah, err := mdl.svc.DoctorIjazahUpload(doctor.DoctorIjazahDataModel{DoctorIjazah: formIjazah})
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Failed to upload Ijazah", nil))
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
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to update doctor data", nil))
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
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to get doctor id", nil))
		}

		role := mdl.jwt.CheckRole(c)
		fmt.Println(role)

		if role != "Doctor" && role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Insufficient permissions", nil))
		}

		var input = new([]DoctorExperience)

		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler: Bind Input Error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to bind request", nil))
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
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to update doctor experience", nil))
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
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to get doctor id", nil))
		}

		role := mdl.jwt.CheckRole(c)
		fmt.Println(role)

		if role == "Patient" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Insufficient permissions", nil))
		}

		var input = new([]DoctorEducation)
		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler: Bind Input Error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to bind request", nil))
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
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to update doctor education", nil))
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
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to get doctor id", nil))
		}

		role := mdl.jwt.CheckRole(c)
		fmt.Println(role)

		if role == "Patient" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Insufficient permissions", nil))
		}

		var input = new([]DoctorWorkdays)
		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler: Bind Input Error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to bind input request", nil))
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
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to update doctor workadays", nil))
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
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to get doctor rating id", nil))
		}

		getIDPatient, err := mdl.jwt.GetID(c)

		if err != nil {
			c.Logger().Error("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to get patient id", nil))
		}

		role := mdl.jwt.CheckRole(c)
		fmt.Println(role)

		if role != "Patient" && role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Insufficient permissions", nil))
		}

		var input = new(DoctorRating)
		if err := c.Bind(input); err != nil {
			c.Logger().Error("Handler: Bind Input Error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to bind doctor rating request", nil))
		}

		var serviceInput = new(doctor.DoctorRating)

		serviceInput.DoctorReview = input.DoctorReview
		serviceInput.DoctorStarRating = input.DoctorStarRating

		result, err := mdl.svc.UpdateDoctorRating(id, int(getIDPatient), *serviceInput)

		if err != nil {
			c.Logger().Error("Handler: Update Process Error (UpdateDoctor): ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to give review to doctor", nil))
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
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to get id doctor", nil))
		}

		role := mdl.jwt.CheckRole(c)
		fmt.Println(role)

		if role != "Doctor" && role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Insufficient permissions", nil))
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
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to get id doctor", nil))
		}

		role := mdl.jwt.CheckRole(c)
		fmt.Println(role)

		if role != "Doctor" && role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Insufficient permissions", nil))
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

		return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail, path param type or path param id not found", nil))
	}
}

func (mdl *DoctorHandler) InsertDataDoctor() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramType = c.Param("type")

		role := mdl.jwt.CheckRole(c)
		fmt.Println(role)

		if role != "Doctor" && role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Insufficient permissions", nil))
		}

		if paramType == "workday" {

			var input = new(DoctorWorkdays)
			if err := c.Bind(input); err != nil {
				c.Logger().Info("Handler: Bind Input Error: ", err.Error())
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to bind request doctor workadays", nil))
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
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to bind request doctor education", nil))
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
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to bind request doctor experience", nil))
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

func (mdl *DoctorHandler) DoctorDashboard() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mdl.jwt.CheckRole(c)
		if role != "Doctor" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized permission", nil))
		}
		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Error("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to get id doctor", nil))
		}

		res, err := mdl.svc.DoctorDashboard(id)

		if err != nil {
			c.Logger().Error("Handler: Callback process error: ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		var response = new(DashboardResponse)
		response.TotalPatient = res.TotalPatient
		response.TotalJamPraktek = res.TotalJamPraktek
		response.TotalLayananChat = res.TotalLayananChat
		response.TotalLayananVideoCall = res.TotalLayananVideoCall

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get dashboard doctor", response))
	}
}

func (mdl *DoctorHandler) DoctorDashboardPatient() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mdl.jwt.CheckRole(c)
		if role != "Doctor" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized permission", nil))
		}

		var paramID = c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Error("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Fail to get id doctor", nil))
		}
		res, err := mdl.svc.DoctorDashboardPatient(id)
		if err != nil {
			c.Logger().Error("Handler: Callback process error: ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success", res))
	}
}

func (mdl *DoctorHandler) DoctorDashboardAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mdl.jwt.CheckRole(c)
		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized permission", nil))
		}

		res, err := mdl.svc.DoctorDashboardAdmin()

		if err != nil {
			c.Logger().Error("Handler: Doctor Dashboard Process Error: ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Doctor Dashboard Process Error", nil))
		}

		var response = new(DashboardAdminResponse)
		response.TotalDoctor = res.TotalDoctor
		response.TotalDoctorActive = res.TotalDoctorActive
		response.TotalDoctorBaru = res.TotalDoctorBaru
		response.TotalDoctorPending = res.TotalDoctorPending

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Get Doctor Dashboard", response))
	}
}

func (mdl *DoctorHandler) DenyDoctor() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mdl.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}
		var paramID = c.Param("user_id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			c.Logger().Error("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid User Input Param ID", nil))
		}

		_, err = mdl.svc.DenyDoctor(id)

		if err != nil {
			c.Logger().Error("Handler : Deny Doctor Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Deny Doctor Process Failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Denny Doctor", nil))
	}
}

func (mdl *DoctorHandler) ApproveDoctor() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := mdl.jwt.CheckRole(c)

		if role != "Admin" {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Only admin can access this page", nil))
		}
		var paramID = c.Param("user_id")
		id, err := strconv.Atoi(paramID)

		if err != nil {
			c.Logger().Error("Handler : Param ID Error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid User Input Param ID", nil))
		}

		_, err = mdl.svc.ApproveDoctor(id)

		if err != nil {
			c.Logger().Error("Handler : Approve Doctor Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Approve Doctor Process Failed", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success Approve Doctor", nil))
	}
}
