package io.roost.ecommerce;

import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;

import io.roost.ecommerce.model.Product;
import io.roost.ecommerce.service.ProductService;

@SpringBootApplication
public class EcommerceApplication {

	public static void main(String[] args) {
		SpringApplication.run(EcommerceApplication.class, args);
	}

	@Bean
	CommandLineRunner runner(ProductService productService) {
		return args -> {
			productService.save(new Product(1L, "TV", 300.00, "TV"));
            productService.save(new Product(2L, "Console", 200.00, "Console"));
            productService.save(new Product(3L, "Headphone", 275.00, "Headphone"));
            productService.save(new Product(4L, "Keyboard", 100.00, "Keyboard"));
            productService.save(new Product(5L, "Laptop", 699.00, "Laptop"));
            productService.save(new Product(6L, "Watch", 300.00, "Watch"));
			productService.save(new Product(7L, "FlipPhone", 500.00, "FlipPhone"));
			productService.save(new Product(8L, "Speaker", 350.00, "Speaker"));
			productService.save(new Product(9L, "Tablet", 400.00, "Tablet"));
		};
	}

}
