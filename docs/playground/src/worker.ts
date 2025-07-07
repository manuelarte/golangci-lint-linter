import './wasm_exec';
import {initWASM, YAMLFuncMap, YAMLProcessResult} from './YAML.ts';

const yaml = initWASM();
const funcMap = yaml.then((v): Promise<YAMLFuncMap> => {
    const apply = (code: string): Promise<YAMLProcessResult> => {
        const res = v.apply(code);
        return new Promise((resolve) => {
            resolve({
                result: res,
            })
            return
        });
    };
    return new Promise((resolve) => {
        resolve({
            apply: apply,
        });
    })
});

self.addEventListener('message', (e) => {
    const data = e.data as {
        code: string
    }
    funcMap.then((v) => {
        v.apply(data.code).then((value) => {
            self.postMessage(value);
        })
    })
});

export default {}
