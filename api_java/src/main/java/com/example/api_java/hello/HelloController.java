package com.example.api_java.hello;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class HelloController {

    @GetMapping("/hi")
    public HelloResponse sayHi() {
        return new HelloResponse("Hello world api with Java");
    }

}
