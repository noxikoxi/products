import {Link} from "react-router-dom";

type Props = {
    links: string[],
    labels: string[]
}

const Navbar = ({links, labels} : Props) => {

    return (
    <nav className="bar">
        <ul>
            {links.map((link, i) => (
                <li key={link} style={{fontSize: "1.5em"}}>
                    <Link to={link}>{labels[i]}</Link>
                </li>
            ))}
        </ul>
    </nav>
    )

}

export default Navbar;