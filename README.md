# tod2rgb

Time of Day to RGB (And brightness eventually)

This is an experiment to convert time of day using sunrise/sunset and astronomical noon to determine the RGB value of daylight for things like automation of daylight cycles for aquariums or plants or just plain old LED llghts using WLED.

TODO: add some design detail and references.

The core kelvin to rgb algorithm is adapted from
https://gist.github.com/paulkaplan/5184275 which itseelf was derived from
http://www.tannerhelland.com/4435/convert-temperature-rgb-algorithm-code/