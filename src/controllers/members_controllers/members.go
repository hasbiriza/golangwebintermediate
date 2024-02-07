package members_controller

import (
	"api_tugas_minggu4/src/helper"
	"api_tugas_minggu4/src/middleware"
	"api_tugas_minggu4/src/models/members_models"
	models "api_tugas_minggu4/src/models/members_models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// ////////////////////////////////////////////////////Register&Login////////////////////////////////////
func RegisterSeller(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "POST" {
		var input members_models.Member
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Invalid request body")
			return
		}

		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		Password := string(hashPassword)

		item := members_models.Member{
			Member_name: input.Member_name,
			Email:       input.Email,
			Password:    Password,
			Role:        "Seller",
			Address:     input.Address,
			Phone:       input.Phone,
		}
		members_models.Create_member(&item)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Seller Registered",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Ke Json", http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(res); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}

}

func RegisterCustomer(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "POST" {
		var input members_models.Member
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request body")
			return
		}
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		Password := string(hashedPassword)

		item := members_models.Member{
			Member_name: input.Member_name,
			Email:       input.Email,
			Password:    Password,
			Role:        "Customer",
			Address:     input.Address,
			Phone:       input.Phone,
		}
		members_models.Create_member(&item)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Customer Registered",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Ke Json", http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(res); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}

	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var input models.Member
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "invalid request body")
			return
		}
		ValidateEmail := models.FindEmail(&input)
		if len(ValidateEmail) == 0 {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "Email is not Found")
			return
		}
		var passwordSecond string
		for _, member := range ValidateEmail {
			passwordSecond = member.Password
		}
		if err := bcrypt.CompareHashAndPassword([]byte(passwordSecond), []byte(input.Password)); err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Password not Found")
			return
		}
		jwtKey := os.Getenv("SECRETKEY")
		token, _ := helper.GenerateToken(jwtKey, input.Email)
		item := map[string]string{
			"Email": input.Email,
			"Token": token,
		}
		res, _ := json.Marshal(item)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	} else {
		http.Error(w, "", http.StatusBadRequest)
	}
}

//////////////////////////////////////////////////CRUD////////////////////////////////////////////

func Data_all_member(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "GET" {
		res, err := json.Marshal(models.SelectAll_member().Value)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(res) // Perubahan disini
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	} else if r.Method == "POST" {
		var member models.Member
		err := json.NewDecoder(r.Body).Decode(&member)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Fprintf(w, "Gagal Decode")
			return
		}

		models.Create_member(&member)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Product Created",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(res) // Perubahan disini
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	} else {
		http.Error(w, "method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func Data_member(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)

	id := r.URL.Path[len("/member/"):]

	if r.Method == "GET" {
		res, err := json.Marshal(models.Select_member(id).Value)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(res) // Perubahan disini
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	} else if r.Method == "PUT" {
		var updateProduct models.Member
		err := json.NewDecoder(r.Body).Decode(&updateProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Fprintf(w, "Gagal Decode boss")
			return
		}

		newProduct := models.Member{
			Member_name: updateProduct.Member_name,
			Email:       updateProduct.Email,
			Password:    updateProduct.Password,
			Role:        updateProduct.Role,
			Address:     updateProduct.Address,
			Phone:       updateProduct.Phone,
		}

		models.Updates_member_seller(id, &newProduct)
		msg := map[string]string{
			"Message": "Member Updated",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(res) // Perubahan disini
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	} else if r.Method == "DELETE" {
		models.Deletes_members(id)
		msg := map[string]string{
			"Message": "Member Deleted",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(res) // Perubahan disini
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	} else {
		http.Error(w, "method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}
