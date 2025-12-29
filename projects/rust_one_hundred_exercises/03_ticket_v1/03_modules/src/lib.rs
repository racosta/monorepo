mod helpers {
  use super::Ticket;

  #[allow(dead_code)]
  pub fn create_todo_ticket(title: String, description: String) -> Ticket {
    Ticket::new(title, description, "To-Do".into())
  }
}

#[allow(dead_code)]
struct Ticket {
  title: String,
  description: String,
  status: String,
}

impl Ticket {
  #[allow(dead_code)]
  fn new(title: String, description: String, status: String) -> Ticket {
    if title.is_empty() {
      panic!("Title cannot be empty");
    }
    if title.len() > 50 {
      panic!("Title cannot be longer than 50 bytes");
    }

    if description.is_empty() {
      panic!("Description cannot be empty");
    }
    if description.len() > 500 {
      panic!("Description cannot be longer than 500 bytes");
    }

    let allowed_statuses = ["To-Do", "In Progress", "Done"];
    if !allowed_statuses.contains(&status.as_str()) {
      panic!("Only `To-Do`, `In Progress`, and `Done` statuses are allowed");
    }

    Ticket {
      title,
      description,
      status,
    }
  }
}

#[cfg(test)]
mod tests {
  use super::*;
  use common::{overly_long_description, overly_long_title, valid_description, valid_title};

  // Use same tests as in the validation module

  #[test]
  #[should_panic(expected = "Title cannot be empty")]
  fn title_cannot_be_empty() {
    Ticket::new("".into(), valid_description(), "To-Do".into());
  }

  #[test]
  #[should_panic(expected = "Description cannot be empty")]
  fn description_cannot_be_empty() {
    Ticket::new(valid_title(), "".into(), "To-Do".into());
  }

  #[test]
  #[should_panic(expected = "Title cannot be longer than 50 bytes")]
  fn title_cannot_be_longer_than_fifty_chars() {
    Ticket::new(overly_long_title(), valid_description(), "To-Do".into());
  }

  #[test]
  #[should_panic(expected = "Description cannot be longer than 500 bytes")]
  fn description_cannot_be_longer_than_500_chars() {
    Ticket::new(valid_title(), overly_long_description(), "To-Do".into());
  }

  #[test]
  #[should_panic(expected = "Only `To-Do`, `In Progress`, and `Done` statuses are allowed")]
  fn status_must_be_valid() {
    Ticket::new(valid_title(), valid_description(), "Funny".into());
  }

  #[test]
  fn done_is_allowed() {
    Ticket::new(valid_title(), valid_description(), "Done".into());
  }

  #[test]
  fn in_progress_is_allowed() {
    Ticket::new(valid_title(), valid_description(), "In Progress".into());
  }

  #[test]
  fn to_do_helper_is_allowed() {
    helpers::create_todo_ticket(valid_title(), valid_description());
  }
}
