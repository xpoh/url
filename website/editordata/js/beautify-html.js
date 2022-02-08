function style_html(e,a,c,d){var g,f;function g(){this.pos=0;this.token="";this.current_mode="CONTENT";this.tags={parent:"parent1",parentcount:1,parent1:""};this.tag_type="";this.token_text=this.last_token=this.last_text=this.token_type="";this.Utils={whitespace:"\n\r\t ".split(""),single_token:"br,input,link,meta,!doctype,basefont,base,area,hr,wbr,param,img,isindex,?xml,embed,script".split(","),extra_liners:"head,body,/html".split(","),in_array:function(k,h){for(var j=0;j<h.length;j++){if(k===h[j]){return true}}return false}};this.get_content=function(){var h="";var k=[];var l=false;while(this.input.charAt(this.pos)!=="<"){if(this.pos>=this.input.length){return k.length?k.join(""):["","TK_EOF"]}h=this.input.charAt(this.pos);this.pos++;this.line_char_count++;if(this.Utils.in_array(h,this.Utils.whitespace)){if(k.length){l=true}this.line_char_count--;continue}else{if(l){if(this.line_char_count>=this.max_char){k.push("\n");for(var j=0;j<this.indent_level;j++){k.push(this.indent_string)}this.line_char_count=0}else{k.push(" ");this.line_char_count++}l=false}}k.push(h)}return k.length?k.join(""):""};this.get_script=function(){var h="";var j=[];var k=new RegExp("<\/script>","igm");k.lastIndex=this.pos;var i=k.exec(this.input);var l=i?i.index:this.input.length;while(this.pos<l){if(this.pos>=this.input.length){return j.length?j.join(""):["","TK_EOF"]}h=this.input.charAt(this.pos);this.pos++;j.push(h)}return j.length?j.join(""):""};this.record_tag=function(h){if(this.tags[h+"count"]){this.tags[h+"count"]++;this.tags[h+this.tags[h+"count"]]=this.indent_level}else{this.tags[h+"count"]=1;this.tags[h+this.tags[h+"count"]]=this.indent_level}this.tags[h+this.tags[h+"count"]+"parent"]=this.tags.parent;this.tags.parent=h+this.tags[h+"count"]};this.retrieve_tag=function(h){if(this.tags[h+"count"]){var i=this.tags.parent;while(i){if(h+this.tags[h+"count"]===i){break}i=this.tags[i+"parent"]}if(i){this.indent_level=this.tags[h+this.tags[h+"count"]];this.tags.parent=this.tags[i+"parent"]}delete this.tags[h+this.tags[h+"count"]+"parent"];delete this.tags[h+this.tags[h+"count"]];if(this.tags[h+"count"]==1){delete this.tags[h+"count"]}else{this.tags[h+"count"]--}}};this.get_tag=function(){var h="";var j=[];var l=false;do{if(this.pos>=this.input.length){return j.length?j.join(""):["","TK_EOF"]}h=this.input.charAt(this.pos);this.pos++;this.line_char_count++;if(this.Utils.in_array(h,this.Utils.whitespace)){l=true;this.line_char_count--;continue}if(h==="'"||h==='"'){if(!j[1]||j[1]!=="!"){h+=this.get_unformatted(h);l=true}}if(h==="="){l=false}if(j.length&&j[j.length-1]!=="="&&h!==">"&&l){if(this.line_char_count>=this.max_char){this.print_newline(false,j);this.line_char_count=0}else{j.push(" ");this.line_char_count++}l=false}j.push(h)}while(h!==">");var m=j.join("");var k;if(m.indexOf(" ")!=-1){k=m.indexOf(" ")}else{k=m.indexOf(">")}var i=m.substring(1,k).toLowerCase();if(m.charAt(m.length-2)==="/"||this.Utils.in_array(i,this.Utils.single_token)){this.tag_type="SINGLE"}else{if(i==="script"){this.record_tag(i);this.tag_type="SCRIPT"}else{if(i==="style"){this.record_tag(i);this.tag_type="STYLE"}else{if(i==="a"){var n=this.get_unformatted("</a>",m);j.push(n);this.tag_type="SINGLE"}else{if(i==="textarea"){var n=this.get_unformatted("</textarea>",m);j.push(n);this.tag_type="SINGLE"}else{if(i==="span"){var n=this.get_unformatted("</span>",m);j.push(n);this.tag_type="SINGLE"}else{if(i==="label"){var n=this.get_unformatted("</label>",m);j.push(n);this.tag_type="SINGLE"}else{if(i.charAt(0)==="!"){if(i.indexOf("[if")!=-1){if(m.indexOf("!IE")!=-1){var n=this.get_unformatted("-->",m);j.push(n)}this.tag_type="START"}else{if(i.indexOf("[endif")!=-1){this.tag_type="END";this.unindent()}else{if(i.indexOf("[cdata[")!=-1){var n=this.get_unformatted("]]>",m);j.push(n);this.tag_type="SINGLE"}else{var n=this.get_unformatted("-->",m);j.push(n);this.tag_type="SINGLE"}}}}else{if(i.charAt(0)==="/"){this.retrieve_tag(i.substring(1));this.tag_type="END"}else{this.record_tag(i);this.tag_type="START"}if(this.Utils.in_array(i,this.Utils.extra_liners)){this.print_newline(true,this.output)}}}}}}}}}return j.join("")};this.get_unformatted=function(j,l){if(l&&l.indexOf(j)!=-1){return""}var h="";var m="";var n=true;do{if(this.pos>=this.input.length){return m}h=this.input.charAt(this.pos);this.pos++;if(this.Utils.in_array(h,this.Utils.whitespace)){if(!n){this.line_char_count--;continue}if(h==="\n"||h==="\r"){m+="\n";for(var k=0;k<this.indent_level;k++){m+=this.indent_string}n=false;this.line_char_count=0;continue}}m+=h;this.line_char_count++;n=true}while(m.indexOf(j)==-1);return m};this.get_token=function(){var i;if(this.last_token==="TK_TAG_SCRIPT"){var h=this.get_script();if(typeof h!=="string"){return h}i=js_beautify(h,{indent_size:this.indent_size,indent_char:this.indent_character,indent_level:this.indent_level});return[i,"TK_CONTENT"]}if(this.current_mode==="CONTENT"){i=this.get_content();if(typeof i!=="string"){return i}else{return[i,"TK_CONTENT"]}}if(this.current_mode==="TAG"){i=this.get_tag();if(typeof i!=="string"){return i}else{var j="TK_TAG_"+this.tag_type;return[i,j]}}};this.printer=function(l,k,h,m){this.input=l||"";this.output=[];this.indent_character=k||" ";this.indent_string="";this.indent_size=h||2;this.indent_level=0;this.max_char=m||70;this.line_char_count=0;for(var j=0;j<this.indent_size;j++){this.indent_string+=this.indent_character}this.print_newline=function(p,n){this.line_char_count=0;if(!n||!n.length){return}if(!p){while(this.Utils.in_array(n[n.length-1],this.Utils.whitespace)){n.pop()}}n.push("\n");for(var o=0;o<this.indent_level;o++){n.push(this.indent_string)}};this.print_token=function(i){this.output.push(i)};this.indent=function(){this.indent_level++};this.unindent=function(){if(this.indent_level>0){this.indent_level--}}};return this}f=new g();f.printer(e,c,a);while(true){var b=f.get_token();f.token_text=b[0];f.token_type=b[1];if(f.token_type==="TK_EOF"){break}switch(f.token_type){case"TK_TAG_START":case"TK_TAG_SCRIPT":case"TK_TAG_STYLE":f.print_newline(false,f.output);f.print_token(f.token_text);f.indent();f.current_mode="CONTENT";break;case"TK_TAG_END":f.print_newline(true,f.output);f.print_token(f.token_text);f.current_mode="CONTENT";break;case"TK_TAG_SINGLE":f.print_newline(false,f.output);f.print_token(f.token_text);f.current_mode="CONTENT";break;case"TK_CONTENT":if(f.token_text!==""){f.print_newline(false,f.output);f.print_token(f.token_text)}f.current_mode="TAG";break}f.last_token=f.token_type;f.last_text=f.token_text}return f.output.join("")};