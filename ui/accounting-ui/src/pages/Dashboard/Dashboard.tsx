import React from 'react'
import {NavBarItem} from '../../components/NavBar/NavBar'
import {Footer, NavBar, ContentContainer} from '../../components'


const Dashboard = () => {
    const items = [
        {text: "Home", active: true},
        {text: "About", active: false},
        {text: "Services", active: false},
        {text: "Contact", active: false}
    ] as NavBarItem[];
    return (
        <div>
            <NavBar title="Los Jefferson" items={items}/>
            <ContentContainer>
                <div className="row">
                    <div className="col-md-8 mb-5">
                        <h2>What We Do</h2>
                        <hr/>
                        <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. A deserunt neque tempore
                            recusandae
                            animi soluta quasi? Asperiores rem dolore eaque vel, porro, soluta unde debitis aliquam
                            laboriosam. Repellat explicabo, maiores!</p>
                        <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Omnis optio neque consectetur
                            consequatur magni in nisi, natus beatae quidem quam odit commodi ducimus totam eum,
                            alias,
                            adipisci nesciunt voluptate. Voluptatum.</p>
                        <a className="btn btn-primary btn-lg" href="/">Call to Action &raquo;</a>
                    </div>
                    <div className="col-md-4 mb-5">
                        <h2>Contact Us</h2>
                        <hr/>
                        <address>
                            <strong>Start Bootstrap</strong>
                            <br/>3481 Melrose Place
                            <br/>Beverly Hills, CA 90210
                            <br/>
                        </address>
                        <address>
                            <abbr title="Phone">P:</abbr>
                            (123) 456-7890
                            <br/>
                            <abbr title="Email">E:</abbr>
                            <a href="mailto:#">name@example.com</a>
                        </address>
                    </div>
                    <div className="row">
                        <div className="col-md-4 mb-5">
                            <div className="card h-100">
                                <img className="card-img-top" src="http://placehold.it/300x200" alt=""/>
                                <div className="card-body">
                                    <h4 className="card-title">Card title</h4>
                                    <p className="card-text">Lorem ipsum dolor sit amet, consectetur adipisicing
                                        elit.
                                        Sapiente esse necessitatibus neque sequi doloribus.</p>
                                </div>
                                <div className="card-footer">
                                    <a href="/" className="btn btn-primary">Find Out More!</a>
                                </div>
                            </div>
                        </div>
                        <div className="col-md-4 mb-5">
                            <div className="card h-100">
                                <img className="card-img-top" src="http://placehold.it/300x200" alt=""/>
                                <div className="card-body">
                                    <h4 className="card-title">Card title</h4>
                                    <p className="card-text">Lorem ipsum dolor sit amet, consectetur adipisicing
                                        elit.
                                        Sapiente esse necessitatibus neque sequi doloribus totam ut praesentium
                                        aut.</p>
                                </div>
                                <div className="card-footer">
                                    <a href="/" className="btn btn-primary">Find Out More!</a>
                                </div>
                            </div>
                        </div>
                        <div className="col-md-4 mb-5">
                            <div className="card h-100">
                                <img className="card-img-top" src="http://placehold.it/300x200" alt=""/>
                                <div className="card-body">
                                    <h4 className="card-title">Card title</h4>
                                    <p className="card-text">Lorem ipsum dolor sit amet, consectetur adipisicing
                                        elit.
                                        Sapiente esse necessitatibus neque.</p>
                                </div>
                                <div className="card-footer">
                                    <a href="/" className="btn btn-primary">Find Out More!</a>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </ContentContainer>
            <Footer>
                <p className="m-0 text-center text-white">Copyright &copy; Your Website 2019</p>
            </Footer>
        </div>
    );
}

export default Dashboard;