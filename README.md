# Clean-Architecture-Based-Microservices

1. Project Overview and Purpose  
   This system is a distributed platform for managing healthcare interactions. It is split into two primary services: Doctor Service and Appointment Service. With separated services it is easier to scale them independently, and each service uses its own database, so schema change in one does not break the other

2. Service Responsibility.  
   Doctor Service: Create Doctors, List Doctors (all and by ID)  
   Appointment Service: Create Appointment, Update status, List appointments (all and by ID)

3. Folder Structure and Dependency Flow.
   Both service follow the Clean Architecture dependency rule.
   - internal/domain/ contains entities and interfaces
   - internal/usecase/ contains business logic, depends only only on Domain
   - internal/delivery/ contains HTTP handlers (Gin) and DTOs, depends only usecase and domain
   - internal/repository/ contains Database logic (MongoDB), depends on domain

4. Inter-Service Communication  
   Appointment service acts as a client to the octor Service during the CreateAppointment via synchronous HTTP GET request (GET /doctors/:id)

5. How to Run project Locally
   1) Start doctor service:
        - cd doctor-service
        - go run cmd/main.go
   2) Start Appointment service:
        - cd appointment-service
        - go run cmd/main.go
   3) Use Postman to test the endpoints and interaction with database

6. Why a shared database was not used  
   Each microservice owns its data in its separate database. If both services shared one DB they would be tightly couples.   
   With such separation, for example, Doctor service can switch to another DB (like PostrgeSQL), and Appointment service can stay on MongoDB and nothing will crash due to Independence

7. Failure scenarios and resilience  
   What happens when Doctor Service is unavailable? The Appointment Service will receive a connection error: DoctorExists() returns false, err, and the Appointment Service returns a 500 Internal Server Error to the user  
   In a production-grade system, I would add: 
      - timeouts - set via context so Appointment service doesn't wait forever; 
      - retries - auto retry request again 2-3 times; 
      - circuit breakers - id the Doctor service fails repeatedly, the "circuit opens," and the Appointment Service stops trying for a while, returning a fast "Service Unavailable" message to protect the system from crashing.
