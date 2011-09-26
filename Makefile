include $(GOROOT)/src/Make.inc

TARG=stemmer

GOFILES=\
  stemmer.go \

	
stemmer: stemmer.$(O)
	$(LD) -L _obj -o stemmer stemmer.$(O)

stemmer.$(O): main.go _obj/stemmer.a
	$(GC) -I_obj -o $@ main.go

$(GOBIN)/stemmer: stemmer
	cp stemmer $@

INSTALLFILES+=$(GOBIN)/stemmer
CLEANFILES+=$(GOBIN)/stemmer
CLEANFILES+=./stemmer

include $(GOROOT)/src/Make.pkg
