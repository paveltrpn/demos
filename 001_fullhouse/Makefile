
CC       := g++-10
CFLAGS   := -Wall -std=c++17 -g
LFLAGS 	 := -lstdc++ -lGL -lGLU -lGLEW -lglfw -ljpeg -ltga

SRCS := src/main.cpp \
		src/mainApp.cpp \
		src/mesh.cpp \
		src/mtrx2.cpp \
		src/mtrx3.cpp \
		src/mtrx4.cpp \
		src/vec2.cpp \
		src/vec3.cpp \
		src/vec4.cpp \
		src/qtnn.cpp \
		src/plane.cpp \
		src/bitmap.cpp \
		src/bitmap_text.cpp

OBJS := main.o \
		mainApp.o \
		mesh.o \
		mtrx2.o \
		mtrx3.o \
		mtrx4.o \
		vec2.o \
		vec3.o \
		vec4.o \
		qtnn.o \
		plane.o \
		bitmap.o \
		bitmap_text.o

#OBJS := $(SRCS:.cpp=.o)

.PHONY: clean all

all: $(OBJS) main

main.o: src/main.cpp
	$(CC) $(CFLAGS) $(INCDIR) -c src/main.cpp

mainApp.o: src/mainApp.cpp src/mainApp.h
	$(CC) $(CFLAGS) $(INCDIR) -c src/mainApp.cpp

mesh.o: src/mesh.cpp
	$(CC) $(CFLAGS) $(INCDIR) -c src/mesh.cpp

mtrx2.o: src/mtrx2.cpp
	$(CC) $(CFLAGS) $(INCDIR) -c src/mtrx2.cpp

mtrx3.o: src/mtrx3.cpp
	$(CC) $(CFLAGS) $(INCDIR) -c src/mtrx3.cpp

mtrx4.o: src/mtrx4.cpp
	$(CC) $(CFLAGS) $(INCDIR) -c src/mtrx4.cpp

vec2.o: src/vec2.cpp
	$(CC) $(CFLAGS) $(INCDIR) -c src/vec2.cpp

vec3.o: src/vec3.cpp
	$(CC) $(CFLAGS) $(INCDIR) -c src/vec3.cpp

vec4.o: src/vec4.cpp
	$(CC) $(CFLAGS) $(INCDIR) -c src/vec4.cpp

qtnn.o: src/qtnn.cpp
	$(CC) $(CFLAGS) $(INCDIR) -c src/qtnn.cpp

plane.o: src/plane.cpp
	$(CC) $(CFLAGS) $(INCDIR) -c src/plane.cpp

bitmap.o: src/bitmap.cpp src/bitmap.h
	$(CC) $(CFLAGS) $(INCDIR) -c src/bitmap.cpp

bitmap_text.o: src/bitmap_text.cpp src/bitmap_text.h
	$(CC) $(CFLAGS) $(INCDIR) -c src/bitmap_text.cpp

main: $(OBJS)
	$(CC) $(OBJS) $(LIBDIR) -o main $(LFLAGS)

clean:
	rm -f *.o main
