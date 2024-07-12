div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div md:w-1/2 flex justify-center
      div
        div 
          form id=honey text-center text-2xl space-y-3 mt-3
            div
              text1
            div
              input type=text input input-primary id=text1
            div
              text2
            div
              input type=text input input-primary id=text2
            div
              input id=save type=submit value=Save btn btn-primary btn-sm mt-3
        div 
          {{ range $i, $item := .items }}
            div font-mono
              a href=/core/honey/{{$item.guid}}
                {{ add $i 1 }}.  {{ $item.text1 }}
          {{ end }}
