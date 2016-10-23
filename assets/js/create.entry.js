import Vue from 'vue';
import axios from 'axios';

class Markdown extends Vue {
  constructor() {
    var properties = {
      delimiters: ['${', '}'],
      el: '#md-editor',
      data: {
        snippetTitle: '',
        markdownInputSrc: '',
        parsedMarkdown: '',
        isActiveEditor: true,
        isActivePreview: false,
      },
      methods: {
        markdownPreview: () => {
          this.parseMarkdown();
        },
        dispMarkdownEditor: () => {
          this.switchActiveEditor(true);
        },
        vPostSnippet: () => {
          this.postSnippet();
        }
      }
    };
    super(properties);
  }

  parseMarkdown() {
    axios({
      method: 'post',
      url: '/api/v1/preview',
      headers: {'Content-Type': 'application/json'},
      data: {
        markdown: this.markdownInputSrc
      }
    })
    .then(res => {
      this.parsedMarkdown = res.data.parsedStr;
      this.switchActiveEditor(false);
    })
    .catch(err => {
      console.log(err);
    });
  }

  switchActiveEditor(switchFlg) {
    this.isActiveEditor = switchFlg;
    this.isActivePreview = !switchFlg;
  }

  postSnippet() {
    axios({
      method: 'post',
      url: '/api/v1/snippet/create',
      headers: {'Content-Type': 'application/json'},
      data: {
        snippetTitle: this.snippetTitle,
        snippetText: this.markdownInputSrc
      }
    })
    .then(res => {
      location.href = "/";
    })
    .catch(err => {
      console.log(err);
    });
  }
}


document.addEventListener('DOMContentLoaded', () => {
  let md = new Markdown();
});
