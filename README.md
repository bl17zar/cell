# distsim

distsim is a small experiment project. it demonstrates disturbance propagation.

## features

* arbitrary disturbance seed;
* arbitrary obstacles;
* interference and reflections simulations;
* simulation cycles detection;

## how it works

in the computing heart of the project lies a graph;
**on every step**:

* leaf nodes try to produce descendents. they can't produce them if the place is taken by obstacle or other node;
* we search for cycles and remove them (interference simulation);

<img width="719" alt="Screen Shot 2022-05-30 at 16 01 06" src="https://user-images.githubusercontent.com/32969427/170988224-b19bf8a9-a1c4-412a-be6f-348f2e4ff9c9.png">

# roadmap

* add interface for setting disturbance seed (gocui);
* add interface for settings arbitrary obstacles (gocui);
* wrap drawer into gocui;
