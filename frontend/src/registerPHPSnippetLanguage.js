/**
 * register php-snippet language
 * based on from https://github.com/microsoft/monaco-editor/blob/main/src/basic-languages/php/php.ts
 * and https://github.com/tinkerun/tinkerun/blob/master/packages/tinkerun-renderer/src/utils/registerPHPSnippetLanguage.js
 *
 * @param {monaco.languages} languages
 */
export const registerPHPSnippetLanguage = languages => {
  const languageId = 'php-snippet'

  languages.register({ id: languageId })

  languages.setLanguageConfiguration(languageId, {
    wordPattern:
      /(-?\d*\.\d\w*)|([^\`\~\!\@\#\%\^\&\*\(\)\-\=\+\[\{\]\}\\\|\;\:\'\"\,\.\<\>\/\?\s]+)/g,

    comments: {
      lineComment: '//',
      blockComment: ['/*', '*/'],
    },

    brackets: [
      ['{', '}'],
      ['[', ']'],
      ['(', ')'],
    ],

    autoClosingPairs: [
      { open: '{', close: '}', notIn: ['string'] },
      { open: '[', close: ']', notIn: ['string'] },
      { open: '(', close: ')', notIn: ['string'] },
      { open: '"', close: '"', notIn: ['string'] },
      { open: "'", close: "'", notIn: ['string', 'comment'] },
    ],

    folding: {
      markers: {
        start: new RegExp('^\\s*(#|//)region\\b'),
        end: new RegExp('^\\s*(#|//)endregion\\b')
      },
    },
  })

  languages.setMonarchTokensProvider(languageId, {
    defaultToken: '',
    tokenPostfix: '',

    tokenizer: {
      root: [
        [
          /[a-zA-Z_]\w*/,
          {
            cases: {
              '@phpKeywords': { token: 'keyword.php' },
              '@phpCompileTimeConstants': { token: 'constant.php' },
              '@default': 'identifier.php',
            },
          },
        ],
        [
          /[$a-zA-Z_]\w*/,
          {
            cases: {
              '@phpPreDefinedVariables': {
                token: 'variable.predefined.php',
              },
              '@default': 'variable.php',
            },
          },
        ],

        // brackets
        [/[{}]/, 'delimiter.bracket.php'],
        [/[\[\]]/, 'delimiter.array.php'],
        [/[()]/, 'delimiter.parenthesis.php'],

        // whitespace
        [/[ \t\r\n]+/],

        // comments
        [/(#|\/\/)$/, 'comment.php'],
        [/(#|\/\/)/, 'comment.php', '@phpLineComment'],

        // block comments
        [/\/\*/, 'comment.php', '@phpComment'],

        // strings
        [/"/, 'string.php', '@phpDoubleQuoteString'],
        [/'/, 'string.php', '@phpSingleQuoteString'],

        // delimiters
        [/[\+\-\*\%\&\|\^\~\!\=\<\>\/\?\;\:\.\,\@]/, 'delimiter.php'],

        // numbers
        [/\d*\d+[eE]([\-+]?\d+)?/, 'number.float.php'],
        [/\d*\.\d+([eE][\-+]?\d+)?/, 'number.float.php'],
        [/0[xX][0-9a-fA-F']*[0-9a-fA-F]/, 'number.hex.php'],
        [/0[0-7']*[0-7]/, 'number.octal.php'],
        [/0[bB][0-1']*[0-1]/, 'number.binary.php'],
        [/\d[\d']*/, 'number.php'],
        [/\d/, 'number.php'],
      ],

      phpComment: [
        [/\*\//, 'comment.php', '@pop'],
        [/[^*]+/, 'comment.php'],
        [/./, 'comment.php'],
      ],

      phpLineComment: [
        [/\?>/, { token: '@rematch', next: '@pop' }],
        [/.$/, 'comment.php', '@pop'],
        [/[^?]+$/, 'comment.php', '@pop'],
        [/[^?]+/, 'comment.php'],
        [/./, 'comment.php'],
      ],

      phpDoubleQuoteString: [
        [/[^\\"]+/, 'string.php'],
        [/@escapes/, 'string.escape.php'],
        [/\\./, 'string.escape.invalid.php'],
        [/"/, 'string.php', '@pop'],
      ],

      phpSingleQuoteString: [
        [/[^\\']+/, 'string.php'],
        [/@escapes/, 'string.escape.php'],
        [/\\./, 'string.escape.invalid.php'],
        [/'/, 'string.php', '@pop'],
      ],
    },

    phpKeywords: [
      '__construct',
      'abstract',
      'and',
      'array',
      'as',
      'break',
      'callable',
      'case',
      'catch',
      'cfunction',
      'class',
      'clone',
      'const',
      'continue',
      'declare',
      'default',
      'die',
      'do',
      'echo',
      'else',
      'elseif',
      'empty',
      'enddeclare',
      'endfor',
      'endforeach',
      'endif',
      'endswitch',
      'endwhile',
      'eval',
      'exit',
      'extends',
      'false',
      'final',
      'finally',
      'fn',
      'for',
      'foreach',
      'function',
      'global',
      'goto',
      'if',
      'implements',
      'include',
      'include_once',
      'instanceof',
      'insteadof',
      'interface',
      'isset',
      'list',
      'match',
      'namespace',
      'new',
      'null',
      'object',
      'old_function',
      'or',
      'print',
      'private',
      'protected',
      'public',
      'readonly',
      'require',
      'require_once',
      'resource',
      'return',
      'static',
      'switch',
      'throw',
      'trait',
      'true',
      'try',
      'unset',
      'use',
      'var',
      'while',
      'xor',
      'yield',
    ],

    phpCompileTimeConstants: [
      '__CLASS__',
      '__DIR__',
      '__FILE__',
      '__LINE__',
      '__NAMESPACE__',
      '__METHOD__',
      '__FUNCTION__',
      '__TRAIT__',
    ],

    phpPreDefinedVariables: [
      '$GLOBALS',
      '$_SERVER',
      '$_GET',
      '$_POST',
      '$_FILES',
      '$_REQUEST',
      '$_SESSION',
      '$_ENV',
      '$_COOKIE',
      '$php_errormsg',
      '$HTTP_RAW_POST_DATA',
      '$http_response_header',
      '$argc',
      '$argv',
    ],

    escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,
  })
}
