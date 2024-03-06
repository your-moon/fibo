"use client";


import EditorJS from "@editorjs/editorjs";
import Header from "@editorjs/header";
import React, { useEffect } from "react";
import { EDITOR_JS_TOOLS } from "./tools";

interface EditorProps {
  data: any;
  onChange: (data: any) => void;
  editorBlock: string;
}

const Editor = ({ data, onChange, editorBlock }: EditorProps) => {
  let ref = React.useRef<EditorJS | null>(null);
  // eslint-disable-next-line react-hooks/exhaustive-deps
  const initEditor = () => {
    if (!ref.current) {
      const editor = new EditorJS({
        placeholder: "Let`s write an awesome story!",
        holder: editorBlock,
        tools: EDITOR_JS_TOOLS,
        data: data,
        async onChange(api, _event) {
          const data = await api.saver.save();
          onChange(data);
        },
      });
      ref.current = editor;
    }
  };
  useEffect(() => {
    //Initialize editorjs if we don't have a reference
    if (!ref.current) {
      initEditor();
    }

    //Add a return function to handle cleanup
    return () => {
      if (ref.current && ref.current.destroy) {
        ref.current.destroy();
      }
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <>
      <div>
        <div className="bg-slate-50" id={editorBlock}></div>
      </div>
    </>
  );
};

export default Editor;
