#pragma once

#include <stdint.h>

void CmdHandler();
void CCmdHandler();

int HookedVFadeAlpha();
void HookedRDrawSequentialPoly(uintptr_t surf, int face);
void HookedRClear();
void HookedMemoryInit(uintptr_t buf, int size);
