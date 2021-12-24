import EditorJS from '@editorjs/editorjs'
import Header from '@editorjs/header'
import List from '@editorjs/list'
import Embed from '@editorjs/embed'
import InlineCode from '@editorjs/inline-code'
import Table from '@editorjs/table'

export default (context, inject) => {
  const defaultOptions = {
    id: '',
    data: {},
    onChange: () => {},
  }

  const editor = (options = defaultOptions) => {
    return new EditorJS({
      config: {
        minWidth: 1000,
      },
      placeholder: 'Let`s write an awesome story!',
      /**
       * Id of Element that should contain Editor instance
       */
      holder: options.id,
      /**
       * Available Tools list.
       * Pass Tool's class or Settings object for each Tool you want to use
       */
      tools: {
        header: {
          class: Header,
          inlineToolbar: ['link'],
        },
        list: {
          class: List,
          inlineToolbar: ['link', 'bold'],
        },
        embed: {
          class: Embed,
          inlineToolbar: false,
          config: {
            services: {
              youtube: true,
            },
          },
        },
        inlineCode: {
          class: InlineCode,
          shortcut: 'CMD+SHIFT+M',
        },
        table: Table,
      },
      /**
       * Previously saved data that should be rendered
       */
      data: options.data || {},
      onChange(data) {
        // pass in function from component to run on change
        options.onChange(data)
      },
    })
  }

  inject('editor', (options) => editor(options))
}
