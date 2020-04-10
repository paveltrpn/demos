
#ifndef __tga_h__
#define __tga_h__

#include <iostream>
#include <fstream>
#include <string>
#include <stdio.h>
#include <stdint.h>

using namespace std;

#pragma pack(push, 1)
typedef struct STGAHeader
{
	uint8_t identsize;
	uint8_t colorMapType;
	uint8_t imageType;
	uint16_t colorMapStart;
	uint16_t colorMapLength;
	uint8_t colorMapBits;
	uint16_t xstart;
	uint16_t ystart;
	uint16_t width;
	uint16_t height;
	uint8_t bits;
	uint8_t descriptor;
} STGAHeader;
#pragma pack(pop)

/*
uint8_t *LoadTGAFromFile(const char *file_name,int32_t &width,int32_t &height);
uint8_t *LoadTGAFromResource(HMODULE hModule,int32_t id,int32_t &width,int32_t &height);
bool SaveTGA(const char *file_name,int32_t width,int32_t height,uint8_t *image);
*/
class tga_img_c {
	private:
		void load_dummy();
	
	public:
		int img_width, img_height;
		int img_channels;
		uint8_t *data;

		tga_img_c();
		~tga_img_c();

		bool __from_file(string fname); 
		bool from_file(string name);
		void show_img_stats();
};

#endif
