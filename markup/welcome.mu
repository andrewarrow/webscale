div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2
      div text-center text-2xl
      div mt-9 relative
        <div style="pointer-events: none;" class="border-2 w-[9rem] h-60 z-0 absolute left-0"></div>
        <div style="pointer-events: none;" class="bg-black w-[1rem] h-60 z-0 absolute left-[9rem]"></div>
        <div style="pointer-events: none;" class="border-2 w-[11rem] h-60 z-0 absolute left-[10rem]"></div>
        <div style="pointer-events: none;" class="bg-black w-[1rem] h-60 z-0 absolute left-[21rem]"></div>
        <div style="pointer-events: none;" class="border-2 w-[8.1rem] h-60 z-0 absolute left-[22rem]"></div>
        <video class="z-1 w-[30rem]" controls autoplay muted loop src="https://i.imgur.com/fA99U7n.mp4"/>
      div mt-9 
        a href=https://www.reddit.com/r/webdev/comments/1e1j2ma/comment/lcup5m5 link
          from here
      div mt-9 
        <iframe width="560" height="315" src="https://www.youtube.com/embed/odcy-wzeelE?si=QXd3iewNnqoh6fy7" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
      div mt-9
        <iframe width="560" height="315" src="https://www.youtube.com/embed/Q8pKJZcdke0?si=RpsJIyqfrocChVym" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
