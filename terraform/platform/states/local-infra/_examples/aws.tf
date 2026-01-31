resource "aws_s3_bucket" "pdf" {
  bucket = "notes"
}

resource "aws_iam_user" "client_ui" {
  name = "client-ui"
}

resource "aws_iam_access_key" "client_ui_key" {
  user = aws_iam_user.client_ui.name
}

resource "aws_iam_user" "admin_ui" {
  name = "admin-ui"
}

resource "aws_iam_access_key" "admin_ui_key" {
  user = aws_iam_user.admin_ui.name
}

data "aws_iam_policy_document" "client_ui_policy" {
  statement {
    actions = ["s3:PutObject"]
    resources = [
      aws_s3_bucket.pdf.arn,
      "${aws_s3_bucket.pdf.arn}/incoming/*",
    ]
  }
}

resource "aws_iam_user_policy" "client_ui_policy_attach" {
  user   = aws_iam_user.client_ui.name
  policy = data.aws_iam_policy_document.client_ui_policy.json
}

data "aws_iam_policy_document" "admin_ui_policy" {
  statement {
    actions = [
      "s3:PutObject",
      "s3:GetObject",
      "s3:DeleteObject",
    ]
    resources = [
      aws_s3_bucket.pdf.arn,
      "${aws_s3_bucket.pdf.arn}/*",
    ]
  }
}

resource "aws_iam_user_policy" "admin_ui_policy_attach" {
  user   = aws_iam_user.admin_ui.name
  policy = data.aws_iam_policy_document.admin_ui_policy.json
}


resource "aws_s3_bucket" "filingpapers" {
  bucket        = "clerk-filingpapers"
  force_destroy = true
}
