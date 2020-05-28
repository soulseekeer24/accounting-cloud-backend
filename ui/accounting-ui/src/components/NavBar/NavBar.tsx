import React from 'react'
import '../../pages/Dashboard/styles.css'

export interface NavBarItem {
    text: string;
    active: boolean;

}

interface Props {
    items: NavBarItem[];
    title: string;
}

const NavItem = (props: NavBarItem) => (
    <li className={props.active ? "nav-item" : "nav-item active"}>
        <a className="nav-link" href="/">{props.text}</a>
    </li>);

const NavBar = (props: Props) => {
    return (
        <nav className="navbar navbar-expand-lg navbar-dark bg-dark-custom fixed-top">
            <div className="container">
                <a className="navbar-brand" href="/">{props.title}</a>
                <button className="navbar-toggler" type="button" data-toggle="collapse"
                        data-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false"
                        aria-label="Toggle navigation">
                </button>
                <div className="collapse navbar-collapse" id="navbarResponsive">
                    <ul className="navbar-nav ml-auto">
                        {props.items.map(item => <NavItem text={item.text} active={item.active}/>)}
                    </ul>
                </div>
            </div>
        </nav>

    );
}
export default NavBar;
