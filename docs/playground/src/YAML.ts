import init from './golangci-lint-linter.wasm?init';

export interface YAMLGoFuncMap {
  apply(code: string): string;
}

export interface YAMLFuncMap {
    apply: (code: string) => Promise<YAMLProcessResult>
}

export interface YAMLProcessResult {
    result: string | undefined
}

declare function apply(v: string): string;

export const initWASM = async (): Promise<YAMLGoFuncMap> => {
  const go = new Go();
  return new Promise(resolve => {
    init(
      go.importObject,
    ).then((instance) => {
      go.run(instance);
      resolve({
        apply: apply
      });
    });
  });
};
