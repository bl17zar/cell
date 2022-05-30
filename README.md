# cell
cell is a small experiment project. it demonstrates disturbance propagation in a space with obtacles or without;

## features
* arbitrary disturbance seed;
* arbitrary obstacles;
* interference and reflections simulations;

## how it works
in the computing heart of the project lies a graph;
**on every step**:
* leaf nodes try to produce descendents. they can't produce them if the place is taken by obstacle or other node;
* we search for cycles and remove them (interference simulation);
