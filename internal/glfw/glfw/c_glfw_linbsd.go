// SPDX-License-Identifier: BSD-3-Clause
// SPDX-FileCopyrightText: 2012 The glfw3-go Authors
// SPDX-FileCopyrightText: 2023 The Ebitengine Authors

//go:build freebsd || linux || netbsd || openbsd

package glfw

/*
#ifdef _GLFW_WAYLAND
	#include "glfw/src/wl_init.c"
	#include "glfw/src/wl_monitor.c"
	#include "glfw/src/wl_window.c"
	#include "glfw/src/wayland-idle-inhibit-unstable-v1-client-protocol.c"
	#include "glfw/src/wayland-pointer-constraints-unstable-v1-client-protocol.c"
	#include "glfw/src/wayland-relative-pointer-unstable-v1-client-protocol.c"
	#include "glfw/src/wayland-viewporter-client-protocol.c"
	#include "glfw/src/wayland-xdg-decoration-unstable-v1-client-protocol.c"
	#include "glfw/src/wayland-xdg-shell-client-protocol.c"
#endif
#ifdef _GLFW_X11
	#include "glfw/src/x11_init.c"
	#include "glfw/src/x11_monitor.c"
	#include "glfw/src/x11_window.c"
	#include "glfw/src/glx_context.c"
#endif
#include "glfw/src/xkb_unicode.c"
*/
import "C"
