# Techstack
The ERIA System is not complex, however the implementation and execution require a lot of attention to detail to make sure this system does not fail.

Written below are the specification of the various subsystems that make up the ERIA-lifesystems system life cycle.

___
### Backend

The backend subsystem is built up out of multiple sub processes. These processes include:
* Main ERIA database
* Customer organisation database
* API data handeling and routing

The Databases are written in PostgreSQL. This was chosen by the first development teams for its performance, platform independence, and extensibility.

The backend itself is built in Golang (Go for short). This language was chosen because of its efficiency and simplicity.
___
### Frontend


The frontend is simpler compared to the backend. For the frontend, we chose Flutter, because it is cross-platform compatible, allowing us to release the same software on both iOS and Android.