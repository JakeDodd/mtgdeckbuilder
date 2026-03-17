import "./Button.css"

function Button({ onClick }) {

	const handleClick = () => {
		onClick()
	}

	return (
		<div className="button-container">
			<button className="random-button" onClick={handleClick}>
				Random Card
			</button>
		</div>
	)
}

export default Button
