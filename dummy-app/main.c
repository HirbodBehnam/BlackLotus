#include <stdio.h>
#include <time.h>

int main() {
	FILE *file = fopen("log.txt", "a");
	time_t mytime = time(NULL);
	char *time_str = ctime(&mytime);
	fprintf(file,"Ran at %s\n", time_str);
	fclose(file);
	return 0;
}
