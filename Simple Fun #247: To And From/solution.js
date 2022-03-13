function toAndFrom (a, b, t) {
	const qtd_to_turn_direction = Math.abs(a - b);
	const start_direction_to_walk = a > b ? 'left' : 'right';
	const bigger_number = a > b ? a : b;
	const smallest_number = a > b ? b : a;
	const remaining_steps = t % qtd_to_turn_direction;
	let current_direction = start_direction_to_walk;
	let position = 0;
	const turns = Math.floor(t / qtd_to_turn_direction);
	const need_to_change_direction = turns % 2 === 1;
	const LEFT = 'left';
	const RIGHT = 'right';

    if (need_to_change_direction) {
		current_direction = current_direction === LEFT ? RIGHT : LEFT;
	};

	switch (current_direction) {
		case LEFT:
			position = bigger_number - remaining_steps
			break;

		case RIGHT:
			position = smallest_number + remaining_steps
			break;
	}

	return position;
}

module.exports.toAndFrom = toAndFrom;