CREATE TABLE `questions`
(
    `id`                        bigint(20)          NOT NULL AUTO_INCREMENT,
    `question_text`             text,
    `answers`                   text,
    `explain`                   text,
    `level`                     varchar(30),
    `question_type`             varchar(30),
    `image`                     text,
    `audio`                     text,
    `created_at`                TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`                TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`                TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

INSERT INTO `questions`(`question_text`, `answers`,`explain`, `level`, `question_type`, `image`, `audio`) VALUES (
  'Question 1', 
  '["Answer 1", "Answer 2", "Answer 3", "Answer 4"]', 
  'Explain 1', 
  'Level_A1', 
  'QuestionType_Part1', 
  'https://images.unsplash.com', 
  'https://images.unsplash.com'
),(
  'Question 2', 
  '["Answer 1", "Answer 2", "Answer 3", "Answer 4"]', 
  'Explain 1', 
  'Level_A1', 
  'QuestionType_Part1', 
  'https://images.unsplash.com', 
  'https://images.unsplash.com'
),(
  'Question 3', 
  '["Answer 1", "Answer 2", "Answer 3", "Answer 4"]', 
  'Explain 1', 
  'Level_A1', 
  'QuestionType_Part1', 
  'https://images.unsplash.com', 
  'https://images.unsplash.com'
),(
  'Question 4', 
  '["Answer 1", "Answer 2", "Answer 3", "Answer 4"]', 
  'Explain 1', 
  'Level_A1', 
  'QuestionType_Part1', 
  'https://images.unsplash.com', 
  'https://images.unsplash.com'
),( 
  'Question 5', 
  '["Answer 1", "Answer 2", "Answer 3", "Answer 4"]', 
  'Explain 1', 
  'Level_A1', 
  'QuestionType_Part1', 
  'https://images.unsplash.com', 
  'https://images.unsplash.com'
),(
  'Question 6', 
  '["Answer 1", "Answer 2", "Answer 3", "Answer 4"]', 
  'Explain 1', 
  'Level_A1', 
  'QuestionType_Part1', 
  'https://images.unsplash.com', 
  'https://images.unsplash.com'
),(
  'Question 7', 
  '["Answer 1", "Answer 2", "Answer 3", "Answer 4"]', 
  'Explain 1', 
  'Level_A1', 
  'QuestionType_Part1', 
  'https://images.unsplash.com', 
  'https://images.unsplash.com'
),( 
  'Question 8', 
  '["Answer 1", "Answer 2", "Answer 3", "Answer 4"]', 
  'Explain 1', 
  'Level_A1', 
  'QuestionType_Part1', 
  'https://images.unsplash.com', 
  'https://images.unsplash.com'
),( 
  'Question 9', 
  '["Answer 1", "Answer 2", "Answer 3", "Answer 4"]', 
  'Explain 1', 
  'Level_A1', 
  'QuestionType_Part1', 
  'https://images.unsplash.com', 
  'https://images.unsplash.com'
),( 
  'Question 10', 
  '["Answer 1", "Answer 2", "Answer 3", "Answer 4"]', 
  'Explain 1', 
  'Level_A1', 
  'QuestionType_Part1', 
  'https://images.unsplash.com', 
  'https://images.unsplash.com'
);



