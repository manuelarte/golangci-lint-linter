import {
    AppBar,
    Toolbar,
    Typography,
    Box,
    Stack,
    Button,
    TextField,
    SnackbarCloseReason,
    IconButton,
    Snackbar
} from '@mui/material';
import GitHub from '@mui/icons-material/GitHub';
import ContentCopy from '@mui/icons-material/ContentCopy';
import './App.css'
import React, { useState, useRef, useEffect } from 'react';
import { editor } from 'monaco-editor'
import MonacoEditor, { loader } from '@monaco-editor/react';
import Grid from '@mui/material/Grid2';
import { FitAddon } from '@xterm/addon-fit';
import yamlWorker from "./worker?worker";
import '@xterm/xterm/css/xterm.css';
import { Terminal } from '@xterm/xterm';
import json2mq from 'json2mq';
import useMediaQuery from '@mui/material/useMediaQuery';
import {YAMLProcessResult} from "./YAML.ts";

const themeBlack = '#313131';
const themeWhite = 'white';

const isXS = (): boolean => {
  const isNotXS = useMediaQuery(
    json2mq({
      minWidth: 400,
    })
  );
  return !isNotXS;
};

const contentHeight = (): number => {
  if (isXS()) {
    return 200;
  }
  return 400;
}

const Header = (content: any) => {
  const v = btoa(content.content as string);
  const shareURL = `${window.location.origin}${window.location.pathname}?content=${v}`;
  const [open, setOpen] = useState(false);
  const [shareURLFieldVisibility] = useState('hidden');

  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="fixed">
        <Toolbar sx={{ backgroundColor: themeBlack }}>
          <Typography sx={{ flexGrow: 1, textAlign: 'left' }}>
            manuelarte/golangci-lint-linter Playground
          </Typography>
          <Stack sx={{
            visibility: shareURLFieldVisibility,
            backgroundColor: '#A0A0A0',
            paddingLeft: 0.4,
            paddingTop: 0.2,
          }} direction={'row'} alignItems={'center'}>
            <TextField value={shareURL} sx={{
              backgroundColor: themeWhite,
            }} variant="standard"></TextField>
            <Button size='small' sx={{
              backgroundColor: '#A0A0A0',
              color: themeBlack,
            }}
              onClick={() => {
                navigator.clipboard.writeText(shareURL).then(() => {
                  setOpen(true);
                });
              }}>
              <ContentCopy />
            </Button>
            <Snackbar
              open={open}
              autoHideDuration={1500}
              onClose={
                (_: React.SyntheticEvent | Event,
                  reason?: SnackbarCloseReason) => {
                  if (reason === 'clickaway') {
                    return;
                  }
                  setOpen(false);
                }}
              message="Copied"
            />
          </Stack>
          <GitHubLink></GitHubLink>
        </Toolbar>
      </AppBar>
    </Box>
  )
};

const GitHubLink = () => {
    return (<IconButton href="https://github.com/manuelarte/golangci-lint-linter">
        <GitHub sx={{ color: themeWhite }}></GitHub>
    </IconButton>)
};

const TerminalComponent = (v: any) => {
  const terminalRef = useRef<HTMLDivElement>(null)
  const [terminalInstance, setTerminalInstance] = useState<Terminal | null>(null)
  const fitAddon = new FitAddon();
  const out = v.out;

  useEffect(() => {
    const instance = new Terminal({
      cursorInactiveStyle: 'none',
      cursorStyle: 'bar',
      letterSpacing: 4,
      fontFamily: 'monospace',
      fontSize: 16,
      fontWeightBold: 'bold',
      convertEol: true,
      theme: {
        background: themeBlack,
        cursor: themeBlack,
        brightRed: '#da433a',
      },
    });

    instance.loadAddon(fitAddon);

    if (terminalRef.current) {
      instance.open(terminalRef.current)
    }

    setTerminalInstance(instance)

    return () => {
      instance.dispose()
      setTerminalInstance(null)
    }
  }, [
    terminalRef,
  ]);

  useEffect(() => {
    if (!terminalInstance) {
      return;
    }

    terminalInstance.loadAddon(fitAddon);
    fitAddon.fit();
    terminalInstance.clear();
    terminalInstance.writeln(out);
    const handleResize = () => fitAddon.fit()

    window.addEventListener('resize', handleResize)
    return () => {
      window.removeEventListener('resize', handleResize);
    }
  }, [terminalRef, terminalInstance, out]);

  return (
    <Box component="div" sx={{
      height: contentHeight(),
      width: '100%',
      backgroundColor: themeBlack,
    }}>
      <Box
        component="div"
        ref={terminalRef}
        sx={{
          height: '100%',
          width: '90%',
          textAlign: 'left',
          paddingLeft: 1,
          paddingTop: 1,
        }} />
    </Box>
  )
}

function App() {
  const [content, setContent] = useState<string>('');
  const editorRef = useRef<editor.IStandaloneCodeEditor | null>(null);
  const workerRef = useRef<Worker | null>(null);
  const [out, setOut] = useState<string>('');

  useEffect(() => {
    workerRef.current = new yamlWorker();
    workerRef.current.onmessage = (event) => {
      const v = event.data as YAMLProcessResult;
      setOut(v.result as string);
    }
    return () => {
      workerRef.current?.terminate();
    };
  }, []);
  const onCodeChange = () => {
    if (!editorRef || !editorRef.current) {
      return
    }
    const code = editorRef.current.getValue()!;
    setContent(code);
    if (workerRef.current) {
      workerRef.current.postMessage({
        code: code,
      });
    }
  };
  loader.init().then((monaco) => {
    monaco.editor.defineTheme('go-yaml-theme', {
      base: 'vs-dark',
      inherit: true,
      rules: [],
      colors: {
        'editor.background': themeBlack,
        'editor.selectionHighlightBorder': themeBlack,
        'editor.lineHighlightBackground': themeBlack,
        'editor.selectionBackground': themeBlack,
      },
    });
  });

  const search = window.location.search;
  const yamlDataBinary = new URLSearchParams(search).get('content');

  const onMount = (editor: editor.IStandaloneCodeEditor) => {
    editorRef.current = editor;
    if (yamlDataBinary) {
      const code = atob(yamlDataBinary)
      editor.setValue(code);
      setContent(code);
      if (workerRef.current) {
        workerRef.current.postMessage({
          code: code,
        });
      }
    }
  };

  return (
    <>
      <Header content={content} />
      <Grid container>
        <Grid size={{ xs: 12, md: 6 }}>
          <MonacoEditor
            height={contentHeight()}
            language="yaml"
            theme="go-yaml-theme"
            options={{
              fontSize: 16,
              selectOnLineNumbers: true,
              renderWhitespace: 'all',
              autoIndent: 'none',
            }}
            onChange={onCodeChange}
            onMount={onMount}
          />
        </Grid>
        <Grid size={{ xs: 12, md: 6 }}>
            <TerminalComponent out={out} />
        </Grid>
      </Grid>
    </>
  )
}

export default App
