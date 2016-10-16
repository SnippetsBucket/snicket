import Vue from 'vue';
import axios from 'axios';

class Markdown extends Vue {
  constructor() {
    var properties = {
      delimiters: ['${', '}'],
      el: '#md-editor',
      data: {
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

}


document.addEventListener('DOMContentLoaded', () => {
  let md = new Markdown();
});
