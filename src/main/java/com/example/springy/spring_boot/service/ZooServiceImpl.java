package com.example.springy.spring_boot.service;

import com.example.springy.spring_boot.entity.Animal;
import com.example.springy.spring_boot.repository.AnimalRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Objects;

@Service
public class ZooServiceImpl implements ZooService {

    // ! CONNECTION
    @Autowired
    private AnimalRepository animalRepository;

    @Override
    public Animal saveAnimal(Animal animal) {
        return animalRepository.save(animal);
    }

    @Override
    public List<Animal> fetchAnimalList() {
        return (List<Animal>) animalRepository.findAll();
    }

    @Override
    public Animal fetchAnimal(Long animalId) {
        return animalRepository.findById(animalId).get();
    }

    @Override
    public Animal updateAnimal(Animal animal, Long animalId) {
        Animal animalDB = animalRepository.findById(animalId).get();
        if (Objects.nonNull(animal.getName()) && !"".equalsIgnoreCase(animal.getName())) {
            animalDB.setName(animal.getName());
        }
        if (Objects.nonNull(animal.getAge()) && (animal.getAge()!=0)) {
            animalDB.setAge(animal.getAge());
        }
        if (Objects.nonNull(animal.getType()) && !"".equalsIgnoreCase(animal.getType())) {
            animalDB.setType(animal.getType());
        }

        return animalRepository.save(animalDB);
    }

    @Override
    public void deleteAnimalById(Long animalId) {
        animalRepository.deleteById(animalId);
    }
}
