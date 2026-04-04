declare module '*.module.css' {
  const cssModuleClasses: { readonly [key: string]: string };
  export default cssModuleClasses;
}

declare module '*.module.scss' {
  const scssModuleClasses: { readonly [key: string]: string };
  export default scssModuleClasses;
}

declare module '*.module.sass' {
  const sassModuleClasses: { readonly [key: string]: string };
  export default sassModuleClasses;
}

declare module '*.module.less' {
  const lessModuleClasses: { readonly [key: string]: string };
  export default lessModuleClasses;
}
