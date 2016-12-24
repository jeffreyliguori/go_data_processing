This repo will host code to do data processing for various things (census data, election results, etc.).

This will only host code and not the data itself (places to get the data should live in readmes or have accompanying scripts to fetch them).

The directory structure is as follows:

src: All the libraries
  rawdata_processors: Translates the csvs to useable data types by readers.
  readers: libraries to create usable objects out of the data types by readers.
bin: Executables. These should either be simple wrappers around rawdata_processor libraries or use readers to read data from elsewhere.

Any data should be accompanied by where it comes from; if it doesn't let me know and I'll add it.
