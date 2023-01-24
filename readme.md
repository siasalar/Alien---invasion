Alien Invasion
====

A simulation of an alien invasion using Go.


Requirements
----
Go 1.15 or later
Installation

To install the package, run the following command:

```
go get github.com/siasalar/alien-invasion
```
Usage
---
The invasion package provides a **Run** function takes in a **CityMap**, number of aliens, and a movement threshold as 
parameters. It creates the specified number of aliens and randomly places them on the map. The simulation then runs 
until all aliens have been destroyed or each alien has moved at least the specified number of times. In each iteration, 
each alien moves randomly to one of its neighboring cities. If any aliens end up in the same city, they fight and 
destroy that city.

The map is in a file, with one city per line. The city name is first,
followed by 1-4 directions (north, south, east, or west). Each one represents a
road to another city that lies in that direction.

For example:
```
Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
```

The number of aliens can be provided as a command line argument when running the program.

For example:

```
go run main.go 2
```
This will **start** the simulation with **2** aliens.