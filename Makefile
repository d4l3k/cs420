MDS = $(shell find */ -name "*.md")

PDFS = $(patsubst %.md,%.pdf,$(MDS))

all: $(PDFS)

%.pdf: %.md
	- echo "Processing $<"
	- cd `dirname $<` && pandoc --latex-engine xelatex --include-in-header=../../config-files/rice.tex -o `basename $@` `basename $<`
	- cd `dirname $<` && share `basename $@`

clean:
	- rm $(PDFS)

.PHONY: clean all

