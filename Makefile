MDS = $(shell find */ -name "*.md")
NOTE_PDFS = $(shell find notes -name "*.pdf" | sort)

PDFS = $(patsubst %.md,%.pdf,$(MDS))

all: $(PDFS)

%.pdf: %.md
	- cd `dirname $<` && pandoc --latex-engine xelatex --include-in-header=../../config-files/rice.tex -o `basename $@` `basename $<`
#	- cd `dirname $<` && share `basename $@`

notes.pdf: $(NOTE_PDFS)
	pdfunite $(NOTE_PDFS) notes.pdf

clean:
	- rm $(PDFS)

.PHONY: clean all

