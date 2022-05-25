# gorm-mysql tutorial

This repository contains the work from an exercise I performed
to gain a better understanding of gorm and how it works. I'll 
probably expand upon it to further my understanding of using 
GORM to map golang types to mysql types.

For successful execution of main.go one must first setup a mysql 
instance that is described by the connection string, or modify 
the connection string to accomodate differences in their instance
of mysql.

After this the database go_test_models and it's table go_test_models
must be created in the mysql instance for main.go to properly execute.