package com.example.springy.spring_boot.service;

import com.example.springy.spring_boot.entity.Animal;
import java.util.List;

public interface ZooService {
    Animal saveAnimal(Animal animal);

    List<Animal> fetchAnimalList();
    Animal fetchAnimal(Long animalId);

    Animal updateAnimal(Animal animal, Long animalId);

    void deleteAnimalById(Long animalId);
}
