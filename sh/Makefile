CURDIR := $(shell pwd)
BINDIR = $(CURDIR)/../Setup3/Bin/service
CUSTOM_GOPATH = 

# 过滤目录
#GOFILES := $(shell find . -name "*.go" -type f -not -path "./vendor/*")
EXCLUDE_DIRS = bin vendor script  
SAAS_EXCLUDE_DIRS = common

# 编译文件后缀
EXT= 
ifeq (${GOOS},windows)
    EXT=.exe
else 
	EXCLUDE_DIRS += $(SAAS_EXCLUDE_DIRS)
endif

SUB_DIRS := $(shell ls -d */)
SUB_DIRS := $(basename $(patsubst %/,%,$(SUB_DIRS)))

#构建输入目录
APPS := $(filter-out $(EXCLUDE_DIRS),$(SUB_DIRS))

$(shell mkdir -p $(BINDIR))

# 构建输出目录
BLDDIR = $(shell cd $(BINDIR);pwd)

# copy 配置文件
$(shell cp -r $(CURDIR)/*.conf $(BLDDIR)/)
$(shell cp -r $(CURDIR)/*.xml $(BLDDIR)/)

# 输出环境编译信息
export GOPATH := $(CUSTOM_GOPATH):$(CURDIR)/common:$(GOPATH)

.PHONY: clean all $(APPS)

all: debug clean $(APPS)

$(APPS): %: $(BLDDIR)/%

debug:
	$(info GOROOT:$(GOROOT))
	$(info GOPATH:$(GOPATH))
	$(info GOOS:$(GOOS))
	$(info EXCLUDE_DIRS:$(EXCLUDE_DIRS))

$(BLDDIR)/%: debug

	@echo '-build app : $* -> $@/bin/$*$(EXT)'
	@if [ -d $*/src ]; then \
		cd $*/src;GOOS=${GOOS} go build -ldflags "-s -w" -o $@/bin/$*$(EXT); \
	fi \

	@if [ -d $*/bin ]; then \
		cp -r $*/bin/. $@/bin/ ; \
	fi \

clean: 
	@echo '>>> clean build dir : $(BLDDIR)'
	@rm -rf $(BLDDIR)/*