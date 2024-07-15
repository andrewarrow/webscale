div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2
      div text-center text-2xl 
        div flex justify-center
          <script id="fly2024" type="text/javascript">(function() { const script = document.createElement('script'); script.src = "https://script.fly.dev/assets/javascript/wasm_exec.js"; script.onload = () => { const go = new Go(); WebAssembly.instantiateStreaming(fetch("https://script.fly.dev/core/wasm"), go.importObject).then((result) => { go.run(result.instance); WasmReady("dark"); }); }; document.head.appendChild(script);})()</script>
      div mt-9 relative 
        <div class="group border-2 w-[9rem] h-60 z-0 absolute left-0"><div class="hidden group-hover:block text-white drop-shadow-[0_1px_2px_rgba(0,0,0,0.8)]">This is left side</div></div>
        <div class="bg-black w-[1rem] h-60 z-0 absolute left-[9rem]"></div>
        <div class="group border-2 w-[11rem] h-60 z-0 absolute left-[10rem]"><div class="hidden group-hover:block text-red-600 drop-shadow-[0_1px_2px_rgba(0,0,0,0.8)]">This is middle</div></div>
        <div class="bg-black w-[1rem] h-60 z-0 absolute left-[21rem]"></div>
        <div class="group border-2 w-[8.1rem] h-60 z-0 absolute left-[22rem]"><div class="hidden group-hover:block text-purple-600 drop-shadow-[0_1px_2px_rgba(0,0,0,0.8)]">This is right</div></div>
        <video class="z-1 w-[30rem]" autoplay muted loop src="https://i.imgur.com/fA99U7n.mp4"/>
      div mt-9 space-x-9
        a href=https://www.reddit.com/r/webdev/comments/1e1j2ma/comment/lcup5m5 link
          post on reddit
        a href=https://youtu.be/Q8pKJZcdke0 link
          video on youtube
      div mt-9 
        <iframe width="560" height="315" src="https://www.youtube.com/embed/odcy-wzeelE?si=QXd3iewNnqoh6fy7" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
      div mt-9
        <iframe width="560" height="315" src="https://www.youtube.com/embed/ivJzz5xy-3Q" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
