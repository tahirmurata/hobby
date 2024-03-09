<script setup lang="ts">
</script>

<template>
  <div className="h-screen min-h-screen">
    <div className="navbar bg-base-100">
      <div className="navbar-start">
        <div className="dropdown">
          <div tabindex="0" role="button" className="btn btn-ghost lg:hidden">
            <svg width="24px" height="24px" stroke-width="1.5" viewBox="0 0 24 24" fill="none"
              xmlns="http://www.w3.org/2000/svg" stroke="currentColor">
              <path d="M3 5H21" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
              </path>
              <path d="M3 12H21" stroke="currentColor" stroke-width="1.5" stroke-linecap="round"
                stroke-linejoin="round">
              </path>
              <path d="M3 19H21" stroke="currentColor" stroke-width="1.5" stroke-linecap="round"
                stroke-linejoin="round">
              </path>
            </svg>
          </div>
          <ul tabindex="0" className="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52">
            <li><a href="/#/crash">Crash</a></li>
            <li><a href="/#/dice">Dice</a></li>
            <li><a href="/#/roll">Roll</a></li>
            <li><a href="/#/plinko">Plinko</a></li>
          </ul>
        </div>
        <a className="btn btn-ghost text-xl">Hobby</a>
      </div>
      <div className="navbar-center hidden lg:flex">
        <ul className="menu menu-horizontal px-1">
          <li><a href="/#/crash">Crash</a></li>
          <li><a href="/#/dice">Dice</a></li>
          <li><a href="/#/roll">Roll</a></li>
          <li><a href="/#/plinko">Plinko</a></li>
        </ul>
      </div>
      <div className="navbar-end">
        <a className="btn">Button</a>
      </div>
    </div>
    <div className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2">
      <router-view></router-view>
    </div>
  </div>
</template>
