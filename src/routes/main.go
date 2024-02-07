package routes

import (
	"api_tugas_minggu4/src/controllers/category_controllers"
	members_controller "api_tugas_minggu4/src/controllers/members_controllers"
	"api_tugas_minggu4/src/controllers/order_controllers"
	"api_tugas_minggu4/src/controllers/products_controllers"
	"api_tugas_minggu4/src/controllers/ratings_controllers"
	"api_tugas_minggu4/src/middleware"
	"fmt"
	"net/http"

	"github.com/goddtriffin/helmet"
)

func Router() {
	helmet := helmet.Default()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World ,Ini Adalah Halaman paling awal ya")
	})

	//Rating Section//
	http.Handle("/ratings", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(ratings_controllers.Data_ratings))))
	http.Handle("/rating/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(ratings_controllers.Data_rating))))

	//Order Section//
	http.Handle("/orders/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(order_controllers.Data_orders))))
	http.Handle("/order/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(order_controllers.Data_order))))

	//MemberSection//
	http.Handle("/members", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(members_controller.Data_all_member))))
	http.Handle("/member/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(members_controller.Data_member))))
	/*Login and Register*/
	http.Handle("/login", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(members_controller.Login))))
	http.Handle("/RegisterCustomer", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(members_controller.RegisterCustomer))))
	http.Handle("/RegisterSeller", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(members_controller.RegisterSeller))))

	//ProductSection//
	http.Handle("/product/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(products_controllers.Data_product))))
	http.Handle("/products", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(products_controllers.Data_all_product))))

	/*Pagination , Upload Product , Search Produck */
	http.Handle("/search", http.HandlerFunc(products_controllers.Search_product))
	http.Handle("/upload", http.HandlerFunc(products_controllers.Upload_image_product))
	http.Handle("/product_pagination", middleware.JwtMiddleware(helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(products_controllers.Product_pagination)))))

	//CategorySection//
	http.Handle("/categories", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(category_controllers.Data_categories))))
	http.Handle("/category/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(category_controllers.Data_category))))

}
