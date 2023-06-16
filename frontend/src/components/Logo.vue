<template>
  <div class="logo">
    <slot>
      PAVEL<span>.</span>ONE
      <div class="blink">
        M<span class="flicker-fast">o</span>ni<span class="flicker-slow">t</span>o<span class="flicker-fast2">r</span>
      </div>
    </slot>
  </div>
</template>

<style lang="scss">
.logo {
  font-family: 'Bebas Neue', sans-serif;
  text-transform: uppercase;
  font-size: 7em;
  color: #404040;
  position: relative;
  & > span {
    color: #3ce09d;
  }
  .blink {
    font-size: .3em;
    position: absolute;
    right: 0;
    bottom: .5em;
    color: #3ce09d;
    transform: rotateZ(331deg);
    text-shadow: 1px 1px 15px rgba(#3ce09d, .6);
  }
}

.flicker-slow{
  animation: flicker 3s linear infinite;
}

.flicker-fast{
  animation: flicker 1s linear infinite;
}
.flicker-fast2{
  animation: flicker 1.5s linear infinite;
}

@media (prefers-color-scheme: light) {
  .logo {
    color: #555;
  }
}

@keyframes flicker {
  0%, 19.999%, 22%, 62.999%, 64%, 64.999%, 70%, 100% {
    opacity: .99;

  }
  20%, 21.999%, 63%, 63.999%, 65%, 69.999% {
    opacity: 0.4;
  }
}

</style>

<script>

</script>
