import React from 'react'


interface Props {
    className?: string;
}

const Footer = (props: React.PropsWithChildren<Props>) => (
    <footer className={props.className || "py-5 bg-dark"}>
        <div className="container">
            {props.children}
        </div>
    </footer>
);

export default Footer;