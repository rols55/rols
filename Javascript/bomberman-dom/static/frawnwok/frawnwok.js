// Create Element & shortcuts
export function createNode(type, props, ...children) {
    props = props || {};
    return { type, props, children: [].concat.apply([], children) };
}

export const createElement = (...args) => createNode(...args)

export const node = {
    div: (prorps, ...children)     => { return createNode("div", prorps, ...children) },
    h1: (prorps, ...children)      => { return createNode("h1", prorps, ...children) },
    h2: (prorps, ...children)      => { return createNode("h2", prorps, ...children) },
    h3: (prorps, ...children)      => { return createNode("h3", prorps, ...children) },
    label: (prorps, ...children)   => { return createNode("label", prorps, ...children) },
    ul: (prorps, ...children)      => { return createNode("ul", prorps, ...children) },
    li: (prorps, ...children)      => { return createNode("li", prorps, ...children) },
    span: (prorps, ...children)    => { return createNode("span", prorps, ...children) },
    section: (prorps, ...children) => { return createNode("section", prorps, ...children) },
    header: (prorps, ...children)  => { return createNode("header", prorps, ...children) },
    main: (prorps, ...children)    => { return createNode("main", prorps, ...children) },
    footer: (prorps, ...children)  => { return createNode("footer", prorps, ...children) },
    button: (prorps, ...children)  => { return createNode("button", prorps, ...children) },
    input: (prorps, ...children)   => { return createNode("input", prorps, ...children) },
    a: (prorps, ...children)       => {return createNode("a", prorps, ...children)},
    link: (props, onClick, ...children) => { return createNode("a", { ...props, onClick: (e) => { if (onClick) onClick(e); } }, ...children) },
}


let currentApp = null;
let appRoot = null;

// Initialize App
export function init(root, app) {
    currentApp = app;
    appRoot = document.getElementById(root)
    states = [];
    callCount = -1;
    render(appRoot, ...app());
}

// Rendering
export function render(container, ...vNodes) {
    callCount = -1;
    container.innerHTML = '';
    [...vNodes].forEach(vNode => {
        if (vNode) {
            container.appendChild(_render(vNode))
        }
    });
}

export function _render(vNode) {
    if (!vNode) return;
    if (typeof vNode === 'string') {
        return document.createTextNode(vNode);
    }

    /*
    if (typeof vNode === 'function') {
      return _render(vNode());
    }
    */

    const element = document.createElement(vNode.type);

    for (let prop in vNode.props) {
        if (prop !== 'children') {
            if (prop === 'className' || prop === 'classname' || prop === 'class') {
                element.setAttribute('class', vNode.props[prop])
            } else if (typeof vNode.props[prop] === 'function') {
                element[prop.toLowerCase()] = vNode.props[prop]
            } else if (prop.slice(0,5) === 'data_') {
                element.setAttribute('data-'+prop.slice(5), vNode.props[prop])
            } else {
                element.setAttribute(prop, vNode.props[prop])
            }
        }
    }

    vNode.children.forEach(child => {
        const childElement = _render(child);
        if (childElement) {
            element.appendChild(childElement);
        }
    });

    return element;
}

// State
let callCount = -1
let states = []

export function useState(initValue) {
    const id = ++callCount
    if (states[id]) {
        return states[id]
    }

    const setState = (newValue) => {
        states[id][0] = newValue instanceof Function ? newValue(states[id][0]) : newValue
        callCount = -1;
        render(appRoot, ...currentApp());
    }

    states.push([initValue, setState])
    return [initValue, setState]
}

// Event Listeners
export function useEventListener(props) {
    const { target, event, callbacks } = props;

    function handleEvent(event) {
        callbacks.forEach((cb) => {
            cb(event);
        });
    }

    function connect() {
        target.addEventListener(event, handleEvent);
    }

    function disconnect() {
        target.removeEventListener(event, handleEvent);
    }

    return {
        connect,
        disconnect,
    };
}

// Routing
let routes = {};
export function initRoutes(newRoutes) {
    newRoutes.map(route => routes[route.path] = { node: route.node, container: route.container, params: route.params });
    loadInitialRoute();

    const handleChange = () => {
        console.log('handleChange')
        loadInitialRoute();
    }
    const hashListener = useEventListener({target: window, event: "hashchange", callbacks: [handleChange]})
    hashListener.connect();
}

function loadInitialRoute() {
    const pathname = window.location.hash.slice(1);
    if (pathname) {
        navigateTo(pathname);
    } else {
        navigateTo('/');
    }
}

export function navigateTo(pathname) {
    const route = routes[pathname]
    if (route) {
        render(appRoot, ...currentApp({path: pathname, ...route.params}))
    } else {
        if (appRoot) {
            appRoot.innerHTML = '404 - Page Not Found';
        } else {
            alert('404 - Page Not Found')
            navigateTo('/');
        }
    }
}

export const router = { initRoutes, navigateTo }

