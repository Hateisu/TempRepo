package com.example.springy.spring_boot.controller;

import com.example.springy.spring_boot.entity.Animal;
import com.example.springy.spring_boot.service.ZooService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.net.URI;
import java.net.URISyntaxException;
import java.util.List;
import java.util.NoSuchElementException;


@RestController
public class ZooController {
    private static final String ZooMapping = "/animal";
    private static final String ZooMappingWithId = "/animal/{id}";
    private static final String ZooMappingAll = "/animals";

    @Autowired
    private ZooService zooService;

    @PostMapping(ZooMappingAll)
    public ResponseEntity<Animal> saveAnimal(@RequestBody Animal animal) throws URISyntaxException {
        Animal result = zooService.saveAnimal(animal);
        return ResponseEntity.created(new URI(ZooMappingAll+result.getId())).body(result);
    }
    @GetMapping(ZooMappingAll)
    public List<Animal> getAllAnimals(){
        return zooService.fetchAnimalList();
    }

    @GetMapping(ZooMappingWithId)
    public ResponseEntity<Animal> getAnimal(@PathVariable("id") Long animalId){
        try {
            Animal animal = zooService.fetchAnimal(animalId);
            return new ResponseEntity<Animal>(animal, HttpStatus.OK);
        } catch (NoSuchElementException e) {
            return new ResponseEntity<Animal>(HttpStatus.NOT_FOUND);
        }

    }
    @PutMapping(ZooMappingWithId)
    public ResponseEntity<Animal> updateAnimal(@PathVariable("id") Long animalId, @RequestBody Animal animal){
        try {
            return new ResponseEntity<Animal>(zooService.updateAnimal(animal,animalId),HttpStatus.OK);
        }catch (NoSuchElementException e){
            return new ResponseEntity<Animal>(HttpStatus.NOT_FOUND);
        }
    }
    @DeleteMapping(ZooMappingWithId)
    public String deleteAnimalById(@PathVariable("id") Long animalId){
        zooService.deleteAnimalById(animalId);
        return "Deleted Successfully";
    }
}
