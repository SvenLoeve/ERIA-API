# Tech stack
The ERIA System is not complex, however the implementation and execution require a lot of attention to detail to make sure this system does not fail.

Written below are the specification of the various subsystems that make up the ERIA-lifesystems system life cycle.

___
### Backend

The backend subsystem is built up out of multiple sub processes. These processes include:
* Main ERIA database
* Customer organisation database
* API data handling and routing

The Databases are written in PostgreSQL. This was chosen by the first development teams for its performance, platform independence, and extensibility.

The backend itself is built in Golang (Go for short). This language was chosen because of its efficiency and simplicity.
___

### Frontend

The frontend of the ERIA-Lifesystems lifecycle (currently) consists of 3 separate applications, those are the following:
* The main Ambulance application mainly internally referred to as MedApp.
* The Backoffice program used to observe and log the MedApp.
* The personal application for interaction with the hardware element.

These 3 applications complete the current scope of the lifecycle. In the future More will be needed to increase efficiency and workflow capability within ERIA itself.

The scope for the frontend is and easy to use for the target audience. To adhere to this scope a consistent feedback loop is required and close cooperation between all parties.
A scrum environment helps with this as it enables a consistent update schedule and easy release date estimation for fixes and updates. 
___

## MedApp

The MedApp is built in Flutter (Dart) This language was chosen for its versatility and regular use. this means great code support and troubleshooting.
This application's main purpose is showing the data from a target person. 

Since the application has two different locations of use within either ambulance services and care homes there is the possibility to use 2 different search identifier groups. 
these are based on credentials and or link to organisation.

The following is a set of input identifiers:
* Personal identifier (Target = BSN)
* Name & Surname (Target)
* Home Address (Target)
* Phone number (Target)
* Room number (Target)
* Any specific request from primary stakeholders (based on feedback)
___

## Backoffice

The main use of the Backoffice is to show the use of the MedApp and show detailed information. 