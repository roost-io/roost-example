package io.roost.ecommerce.repository;

import org.springframework.data.repository.CrudRepository;

import io.roost.ecommerce.model.Product;

public interface ProductRepository extends CrudRepository<Product, Long> {
    
}
